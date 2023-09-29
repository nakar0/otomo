package model

const (
	maxMonthlySurplusSendCount = 30
	maxDailySendCount          = 5
)

type MonthlySurplusMessageSentCount struct {
	YearMonth YearMonth
	Daily     DailyMessageSentCountList
}

func NewMonthlySurplusMessageSentCount(
	yearMonth YearMonth,
	daily []*DailyMessageSentCount,
) *MonthlySurplusMessageSentCount {
	return &MonthlySurplusMessageSentCount{
		YearMonth: yearMonth,
		Daily:     daily,
	}
}

func (m *MonthlySurplusMessageSentCount) IsRemaining() bool {
	return m.Count() < maxMonthlySurplusSendCount
}
func (m *MonthlySurplusMessageSentCount) IsRemainingDay(day Day) bool {
	d := m.Daily.WhereByDay(day)
	return d.IsRemaining() || m.IsRemaining()
}

func (m *MonthlySurplusMessageSentCount) Count() int {
	var exceeded int
	for _, d := range m.Daily {
		exceeded += d.CountExceeded()
	}
	return exceeded
}

func (m *MonthlySurplusMessageSentCount) CountRemaining() int {
	remaining := maxMonthlySurplusSendCount - m.Count()
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (m *MonthlySurplusMessageSentCount) IfSent(
	day Day,
) *MonthlySurplusMessageSentCount {
	var newDaily = make(DailyMessageSentCountList, len(m.Daily))
	copy(newDaily, m.Daily)

	dayIndex := newDaily.IndexByDay(day)
	if dayIndex == -1 {
		d := NewDailyMessageSentCount(day, 1)
		newDaily = append(newDaily, d)
	} else {
		d := NewDailyMessageSentCount(day, newDaily[dayIndex].Count+1)
		newDaily[dayIndex] = d
	}

	return NewMonthlySurplusMessageSentCount(m.YearMonth, newDaily)
}

type DailyMessageSentCount struct {
	Day   Day
	Count int
}

func NewDailyMessageSentCount(day Day, count int) *DailyMessageSentCount {
	return &DailyMessageSentCount{
		Day:   day,
		Count: count,
	}
}

func (d *DailyMessageSentCount) IsRemaining() bool {
	return d.Count < maxDailySendCount
}

func (d *DailyMessageSentCount) CountRemaining() int {
	remaining := maxDailySendCount - d.Count
	if remaining < 0 {
		return 0
	}
	return remaining
}

func (d *DailyMessageSentCount) CountExceeded() int {
	exceeded := d.Count - maxDailySendCount
	if exceeded < 0 {
		return 0
	}
	return exceeded
}

type DailyMessageSentCountList []*DailyMessageSentCount

func (list DailyMessageSentCountList) IndexByDay(day Day) int {
	for i, d := range list {
		if d.Day == day {
			return i
		}
	}
	return -1
}

func (list DailyMessageSentCountList) WhereByDay(
	day Day,
) *DailyMessageSentCount {
	index := list.IndexByDay(day)
	if index == -1 {
		return &DailyMessageSentCount{Day: day, Count: 0}
	}
	return list[index]
}

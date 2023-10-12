import 'package:flutter_test/flutter_test.dart';
import 'package:otomo/entities/date.dart';
import 'package:otomo/entities/policies_agreements.dart';
import 'package:otomo/entities/user.dart';

void main() {
  final now = DateTime.now();
  group('PoliciesAgreements class', () {
    group('isAgreed getter', () {
      test('should return true when agreed20231011At is not null', () {
        final agreements = PoliciesAgreements(
          userId: 'id',
          agreed20231011At: now,
        );
        expect(true, agreements.isAgreed);
      });
      test('should return false when agreed20231011At is null', () {
        final agreements = PoliciesAgreements.disagree('id');
        expect(false, agreements.isAgreed);
      });
    });
    group('canAgree method', () {
      test('should return true when user age is 13 or more', () {
        final birthday =
            Date(year: now.year - 13, month: now.month, day: now.day);
        final user = User(id: 'id', birthday: birthday);
        expect(true, PoliciesAgreements.canAgree(user));
      });
    });
    test('should return false when user age is 12 or less', () {
      final birthday =
          Date(year: now.year - 12, month: now.month, day: now.day);
      final user = User(id: 'id', birthday: birthday);
      expect(false, PoliciesAgreements.canAgree(user));
    });
  });
}

import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:otomo/domains/entities/date.dart';
part 'user.freezed.dart';
part 'user.g.dart';

@freezed
class User with _$User {
  const User._();
  @JsonSerializable(explicitToJson: true)
  const factory User({
    required String id,
    required Date birthday,
  }) = _User;

  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}

extension BirthdayEx on Date {
  int get age {
    int result = DateTime.now().year - year;
    if (DateTime.now().month < month ||
        (DateTime.now().month == month && DateTime.now().day < day)) {
      result--;
    }
    return result;
  }
}

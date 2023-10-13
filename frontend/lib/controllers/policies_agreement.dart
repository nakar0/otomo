import 'package:injectable/injectable.dart';
import 'package:otomo/domains/entities/app_exception.dart';
import 'package:otomo/domains/entities/date.dart';
import 'package:otomo/domains/entities/policies_agreements.dart';
import 'package:otomo/domains/entities/user.dart';
import 'package:otomo/repositories/policies_agreements.dart';
import 'package:otomo/repositories/user.dart';

@injectable
class PoliciesAgreementControllerImpl {
  const PoliciesAgreementControllerImpl(
    this._agreementsRepository,
    this._userRepository,
  );

  final PoliciesAgreementsRepositoryImpl _agreementsRepository;
  final UserRepositoryImpl _userRepository;

  Future<PoliciesAgreements> agree(String userId, Date birthday) async {
    if (!PoliciesAgreements.canAgree(birthday)) {
      throw const AppException(
        message: 'Must be 13 years old or older',
        cause: Cause.invalidArgument,
        domain: Domain.user,
        field: Field.birthday,
      );
    }

    final user = User(id: userId, birthday: birthday);
    await _userRepository.save(user);

    final agreements = PoliciesAgreements(
      userId: userId,
      agreed20231011At: DateTime.now(),
    );
    await _agreementsRepository.save(agreements);
    return agreements;
  }

  Future<PoliciesAgreements> getAgreements(String userId) async {
    try {
      return await _agreementsRepository.get(userId);
    } on AppException catch (e) {
      if (e.cause == Cause.notFound) {
        return PoliciesAgreements.disagree(userId);
      }
      rethrow;
    }
  }
}

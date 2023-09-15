import 'dart:async';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:otomo/controllers/boundary/id_token.dart';
import 'package:otomo/tools/logger.dart';

// Being injected configs/injection.dart
// TODO: Add tests
class IdTokenControllerImpl implements IdTokenController {
  IdTokenControllerImpl(this._firebaseAuth) {
    _firebaseAuth.authStateChanges().listen((user) {
      if (user == null) {
        _idToken = null;
        _cancelTimerIf();
      } else {
        _refreshIdToken();
      }
    });
  }

  final FirebaseAuth _firebaseAuth;

  String? _idToken;
  Timer? _refreshTimer;

  @override
  String? get idToken => _idToken;

  @override
  Future<void> refreshIdToken() => _refreshIdToken();

  Future<void> _refreshIdToken() async {
    _cancelTimerIf();

    final user = _firebaseAuth.currentUser;
    if (user == null) {
      _idToken = null;
      return;
    }

    final idTokenResult = await user.getIdTokenResult();
    _idToken = idTokenResult.token;
    logger.debug(_idToken!);
    logger.info('refreshed id token');

    _setRefreshTimer(idTokenResult.expirationTime);
  }

  void _cancelTimerIf() {
    if (_refreshTimer != null && _refreshTimer!.isActive) {
      _refreshTimer!.cancel();
    }
  }

  void _setRefreshTimer(DateTime? expirationTime) {
    logger.info('set refresh id token timer. expirationTime: $expirationTime');

    if (expirationTime == null) {
      logger.warn('expiration period is null. That is not expected.');
      _refreshTimer = Timer(const Duration(seconds: 10), _refreshIdToken);
      return;
    }

    final expirationPeriod = expirationTime.difference(DateTime.now());

    final refreshDuration = expirationPeriod - const Duration(minutes: 1);
    _refreshTimer = Timer(refreshDuration, _refreshIdToken);
  }
}

import 'dart:async';
import 'dart:io';

import 'package:firebase_auth/firebase_auth.dart' as auth;
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:otomo/controllers/firebase.dart';
import 'package:otomo/entities/app_exception.dart';
import 'package:otomo/entities/account.dart';
import 'package:otomo/tools/logger.dart';

const _appleEmailScope = 'email';

class AuthControllerImpl {
  AuthControllerImpl(this._firebaseAuth, this._googleSignIn);

  final auth.FirebaseAuth _firebaseAuth;
  final GoogleSignIn _googleSignIn;

  Stream<Account?> authStateChanges() => _firebaseAuth
      .authStateChanges()
      .map((user) => user == null ? null : _userToAccount(user));

  Future<String?> getIdToken() async {
    try {
      return await _firebaseAuth.currentUser?.getIdToken();
    } on auth.FirebaseAuthException catch (e) {
      if (e.code == FirebaseExceptionCode.noSuchProvider) {
        logger.warn('no such provider. signing out');
        await _firebaseAuth.signOut();
      }
      rethrow;
    }
  }

  Future<Account> signInWithGoogle() async {
    final credential = await _getGoogleAuthCredential();
    return _signInWithCredential(credential);
  }

  Future<Account> signInWithApple() async {
    if (Platform.isAndroid) {
      throw AppException.unknown('Apple sign in is not supported on Android');
    }

    final appleProvider = auth.AppleAuthProvider();
    appleProvider.addScope(_appleEmailScope);

    late final auth.UserCredential credential;
    if (kIsWeb) {
      credential = await _firebaseAuth.signInWithPopup(appleProvider);
    } else {
      credential = await _firebaseAuth.signInWithProvider(appleProvider);
    }
    return _userToAccount(credential.user);
  }

  Future<Account> _signInWithCredential(auth.AuthCredential credential) async {
    final userCred = await _firebaseAuth.signInWithCredential(credential);
    return _userToAccount(userCred.user);
  }

  Future<auth.OAuthCredential> _getGoogleAuthCredential() async {
    final googleUser = await _googleSignIn.signIn();
    final googleAuth = await googleUser?.authentication;
    return auth.GoogleAuthProvider.credential(
      accessToken: googleAuth?.accessToken,
      idToken: googleAuth?.idToken,
    );
  }

  Future<void> signOut() => _firebaseAuth.signOut();

  FutureOr<void> deleteAccount() {
    try {
      return _firebaseAuth.currentUser?.delete();
    } on auth.FirebaseAuthException catch (e) {
      if (e.code == FirebaseExceptionCode.requiresRecentLogin) {
        throw const AppException(
          message: 'Please re-authenticate',
          cause: Cause.requiresRecentLogin,
          domain: Domain.auth,
          field: Field.none,
        );
      }
      rethrow;
    }
  }

  Future<Account> reauthenticate() async {
    final account = _userToAccount(_firebaseAuth.currentUser);
    switch (account.authProvider) {
      case AuthProvider.google:
        final credential = await _getGoogleAuthCredential();
        return _reauthenticateWithCredential(credential);
      case AuthProvider.apple:
        return _reauthenticateWithApple();
      default:
        throw AppException.unknown(
            'Unknown auth provider: ${account.authProvider}');
    }
  }

  Future<Account> _reauthenticateWithCredential(
    auth.AuthCredential credential,
  ) async {
    final userCred = await _firebaseAuth.currentUser
        ?.reauthenticateWithCredential(credential);
    return _userToAccount(userCred?.user);
  }

  Future<Account> _reauthenticateWithApple() async {
    final appleProvider = _getAppleOAuthProvider();

    late final auth.UserCredential? credential;
    if (kIsWeb) {
      credential = await _firebaseAuth.currentUser
          ?.reauthenticateWithPopup(appleProvider);
    } else {
      credential = await _firebaseAuth.currentUser
          ?.reauthenticateWithProvider(appleProvider);
    }
    return _userToAccount(credential?.user);
  }

  Account _userToAccount(auth.User? user) {
    if (user == null) throw AppException.unknown('User is null');
    return Account(
      uid: user.uid,
      email: user.email,
      authProvider: _getAuthProvider(user),
    );
  }

  AuthProvider _getAuthProvider(auth.User user) {
    if (user.providerData.isEmpty) {
      throw AppException.unknown('User has no provider data');
    }

    final provider = user.providerData.first.providerId;
    switch (provider) {
      case 'google.com':
        return AuthProvider.google;
      case 'apple.com':
        return AuthProvider.apple;
      default:
        throw AppException.unknown('Unknown provider: $provider');
    }
  }

  auth.AppleAuthProvider _getAppleOAuthProvider() =>
      auth.AppleAuthProvider().addScope(_appleEmailScope);
}


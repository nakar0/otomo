import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/view_models/account.dart';
import 'package:otomo/views/pages/account_deletion/index.dart';
import 'package:otomo/views/pages/settings/index.dart';
import 'package:otomo/views/pages/home/index.dart';
import 'package:otomo/views/pages/sign_in/index.dart';
import 'package:otomo/views/pages/sign_in_with_email_link/index.dart';

abstract class Routes {
  static const signIn = '/sign_in';
  static const signInWithEmailLink = '$signIn/email';
  static const home = '/home';
  static const settings = '/settings';
  static const accountDeletion = '$settings/account_deletion';
}

final _key = GlobalKey<NavigatorState>();

final List<GoRoute> _notSignedInPages = [
  GoRoute(
    path: Routes.signIn,
    builder: (context, state) => const SignInPage(),
  ),
  GoRoute(
    path: Routes.signInWithEmailLink,
    builder: (context, state) => const SignInWithEmailLinkPage(),
  ),
];

final List<GoRoute> _signedInPages = [
  GoRoute(
    path: Routes.home,
    builder: (context, state) => const HomePage(),
  ),
  GoRoute(
    path: Routes.settings,
    pageBuilder: (context, state) => const MaterialPage(
      fullscreenDialog: true,
      child: SettingsPage(),
    ),
  ),
  GoRoute(
    path: Routes.accountDeletion,
    builder: (context, state) => const AccountDeletionPage(),
  )
];

final routerProvider = Provider((ref) {
  final account = ref.watch(accountProvider);
  if (account == null) {
    return GoRouter(
      navigatorKey: _key,
      initialLocation: Routes.signIn,
      routes: _notSignedInPages,
    );
  }

  return GoRouter(
    navigatorKey: _key,
    initialLocation: Routes.home,
    routes: [
      ..._notSignedInPages,
      ..._signedInPages,
    ],
  );
});

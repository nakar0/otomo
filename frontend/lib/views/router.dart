import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/view_models/user.dart';
import 'package:otomo/views/pages/account_deletion.dart';
import 'package:otomo/views/pages/settings.dart';
import 'package:otomo/views/pages/home/index.dart';
import 'package:otomo/views/pages/sign_in.dart';
import 'package:otomo/views/pages/sign_in_with_email_link.dart';
import 'package:otomo/views/routes.dart';

final _key = GlobalKey<NavigatorState>();

final List<RouteBase> _notSignedInPages = [
  GoRoute(
    path: Routes.signIn,
    builder: (context, state) => const SignInPage(),
  ),
  GoRoute(
    path: Routes.signInWithEmailLink,
    builder: (context, state) => const SignInWithEmailLinkPage(),
  ),
];

final List<RouteBase> _signedInPages = [
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
  ),
  ...$appRoutes,
];

final routerProvider = Provider((ref) {
  final user = ref.watch(accountProvider);
  if (user == null) {
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

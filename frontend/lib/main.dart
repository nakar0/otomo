import 'dart:async';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/configs/app_config.dart';
import 'package:otomo/configs/firebase_options/dev.dart' as dev_firebase_opt;
import 'package:otomo/configs/firebase_options/local.dart'
    as local_firebase_opt;
import 'package:otomo/configs/injection.dart';
import 'package:otomo/controllers/boundary/id_token.dart';
import 'package:otomo/tools/logger.dart';
import 'package:otomo/views/app.dart';
import 'package:otomo/views/utils/error_handling.dart';

void main() async {
  logger.info(appConfig.toString());
  runZonedGuarded(() async {
    await setup();
    runApp(const ProviderScope(child: App()));
  }, onRunZoneGuardedError);
}

Future<void> setup() async {
  WidgetsFlutterBinding.ensureInitialized();
  setupErrorHandling();
  await initializeFirebase();
  configureInjection();

  // For initialize id token
  await getIt<IdTokenController>().setupIdToken();

  if (appConfig.isLocal) {
    getIt<FirebaseFirestore>()
        .useFirestoreEmulator(appConfig.otomoServerHost, 8080);
    await getIt<FirebaseAuth>()
        .useAuthEmulator(appConfig.otomoServerHost, 9099);
  }
}

Future<void> initializeFirebase() async {
  FirebaseOptions options;

  switch (appConfig.flavor) {
    case Flavor.local:
      options = local_firebase_opt.DefaultFirebaseOptions.currentPlatform;
    case Flavor.dev:
      options = dev_firebase_opt.DefaultFirebaseOptions.currentPlatform;
    default:
      throw Exception('Invalid flavor: ${appConfig.flavor}');
  }

  await Firebase.initializeApp(options: options);
}

void setupErrorHandling() {
  FlutterError.onError = (details) {
    logger.error('Caught error at FlutterError.onError');
    FlutterError.dumpErrorToConsole(details);
    showErrorSnackbar(details.exception);
  };
  PlatformDispatcher.instance.onError = (error, stack) {
    logger.error('Platform Error: $error', stackTrace: stack);
    showErrorSnackbar(error);
    return true;
  };
}

void onRunZoneGuardedError(Object error, StackTrace stack) {
  logger.error('Caught error at ZoneGuarded: $error', stackTrace: stack);
  showErrorSnackbar(error);
}

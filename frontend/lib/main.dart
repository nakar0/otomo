import 'dart:async';

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:firebase_crashlytics/firebase_crashlytics.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_native_splash/flutter_native_splash.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/configs/app_config.dart';
import 'package:otomo/configs/firebase_options/dev.dart' as dev_firebase_opt;
import 'package:otomo/configs/firebase_options/local.dart'
    as local_firebase_opt;
import 'package:otomo/configs/firebase_options/prod.dart' as prod_firebase_opt;
import 'package:otomo/configs/injection.dart';
import 'package:otomo/development/views/development_app.dart';
import 'package:otomo/tools/app_package_info.dart';
import 'package:otomo/tools/logger.dart';
import 'package:otomo/views/app.dart';
import 'package:otomo/views/utils/error_handling.dart';
import 'package:otomo/views/utils/flutter.dart';

void main() async {
  runZonedGuarded(() async {
    await setup();

    // Hide keyboard when app is hot reloaded
    FlutterUtils.hideKeyboard();

    runApp(ProviderScope(
      child: appConfig.isPlayground ? const PlaygroundApp() : const App(),
    ));
  }, onRunZoneGuardedError);
}

Future<void> setup() async {
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  FlutterNativeSplash.preserve(widgetsBinding: widgetsBinding);

  setNewLogger(Logger(level: appConfig.logLevel));

  logger.info(appConfig.toString());
  await AppPackageInfo.init();
  logger.info(AppPackageInfo.expose());

  setupErrorHandling();
  await initializeFirebase();
  await configureInjection();

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
    case Flavor.prod:
      options = prod_firebase_opt.DefaultFirebaseOptions.currentPlatform;
    default:
      throw Exception('Invalid flavor: ${appConfig.flavor}');
  }

  await Firebase.initializeApp(options: options);
}

void setupErrorHandling() {
  FlutterError.onError = (details) {
    logger.error('Caught error at FlutterError.onError');
    FlutterError.dumpErrorToConsole(details);
    FirebaseCrashlytics.instance.recordFlutterFatalError(details);
    if (details.exception is Error) return;
    showErrorSnackbar(details.exception);
  };
  PlatformDispatcher.instance.onError = (error, stack) {
    logger.error('Platform Error: $error', stackTrace: stack);
    FirebaseCrashlytics.instance.recordError(error, stack, fatal: true);
    showErrorSnackbar(error);
    return true;
  };
}

void onRunZoneGuardedError(Object error, StackTrace stack) {
  logger.error('Caught error at ZoneGuarded: $error', stackTrace: stack);
  showErrorSnackbar(error);
}

// File generated by FlutterFire CLI.
// ignore_for_file: lines_longer_than_80_chars, avoid_classes_with_only_static_members
import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

/// Default [FirebaseOptions] for use with your Firebase apps.
///
/// Example:
/// ```dart
/// import 'dev.dart';
/// // ...
/// await Firebase.initializeApp(
///   options: DefaultFirebaseOptions.currentPlatform,
/// );
/// ```
class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        return ios;
      case TargetPlatform.macOS:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for macos - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.windows:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for windows - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.linux:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for linux - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      default:
        throw UnsupportedError(
          'DefaultFirebaseOptions are not supported for this platform.',
        );
    }
  }

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: 'AIzaSyDyrNJlhuViDKaouiwUgaz4scTwIJUvW3M',
    appId: '1:764287781247:web:46860c0f3d16f01f38324a',
    messagingSenderId: '764287781247',
    projectId: 'otomo-dev-396006',
    authDomain: 'otomo-dev-396006.firebaseapp.com',
    storageBucket: 'otomo-dev-396006.appspot.com',
    measurementId: 'G-4VHV76ZFK7',
  );

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyBxl5wAVdAFaDOXGIiF_Rp0eiO6Y5-85gQ',
    appId: '1:764287781247:android:3d61926b1609d8d538324a',
    messagingSenderId: '764287781247',
    projectId: 'otomo-dev-396006',
    storageBucket: 'otomo-dev-396006.appspot.com',
  );

  static const FirebaseOptions ios = FirebaseOptions(
    apiKey: 'AIzaSyDPfVlO8PZ3mXYmtEcGURM2KIbqOHp7YnI',
    appId: '1:764287781247:ios:30e5173bf78622a838324a',
    messagingSenderId: '764287781247',
    projectId: 'otomo-dev-396006',
    storageBucket: 'otomo-dev-396006.appspot.com',
    androidClientId: '764287781247-04jfo3485hq46egsobj35orn463v2lsq.apps.googleusercontent.com',
    iosClientId: '764287781247-gjv61d3k0rl9jhktqkuu44aa6qk9b30j.apps.googleusercontent.com',
    iosBundleId: 'com.nakar0.otomo.develop',
  );
}

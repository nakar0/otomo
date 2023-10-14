// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// InjectableConfigGenerator
// **************************************************************************

// ignore_for_file: unnecessary_lambdas
// ignore_for_file: lines_longer_than_80_chars
// coverage:ignore-file

// ignore_for_file: no_leading_underscores_for_library_prefixes
import 'dart:async' as _i17;

import 'package:cloud_firestore/cloud_firestore.dart' as _i10;
import 'package:firebase_analytics/firebase_analytics.dart' as _i6;
import 'package:firebase_auth/firebase_auth.dart' as _i7;
import 'package:firebase_crashlytics/firebase_crashlytics.dart' as _i8;
import 'package:firebase_dynamic_links/firebase_dynamic_links.dart' as _i9;
import 'package:get_it/get_it.dart' as _i1;
import 'package:grpc/grpc.dart' as _i5;
import 'package:injectable/injectable.dart' as _i2;
import 'package:otomo/configs/injection.dart' as _i23;
import 'package:otomo/controllers/auth.dart' as _i3;
import 'package:otomo/controllers/chat.dart' as _i21;
import 'package:otomo/controllers/location.dart' as _i12;
import 'package:otomo/controllers/place.dart' as _i13;
import 'package:otomo/controllers/policies_agreement.dart' as _i22;
import 'package:otomo/domains/entities/policies_agreements.dart' as _i18;
import 'package:otomo/domains/repo/policies_agreements.dart' as _i14;
import 'package:otomo/domains/repo/user.dart' as _i19;
import 'package:otomo/grpc/generated/chat_service.pbgrpc.dart' as _i4;
import 'package:otomo/grpc/generated/health.pbgrpc.dart' as _i11;
import 'package:otomo/repositories/policies_agreements.dart' as _i15;
import 'package:otomo/repositories/user.dart' as _i20;
import 'package:shared_preferences/shared_preferences.dart' as _i16;

extension GetItInjectableX on _i1.GetIt {
// initializes the registration of main-scope dependencies inside of GetIt
  _i1.GetIt init({
    String? environment,
    _i2.EnvironmentFilter? environmentFilter,
  }) {
    final gh = _i2.GetItHelper(
      this,
      environment,
      environmentFilter,
    );
    final injectableModule = _$InjectableModule();
    gh.singleton<_i3.AuthControllerImpl>(injectableModule.authController);
    gh.singleton<_i4.ChatServiceClient>(injectableModule.chatServiceClient);
    gh.singleton<_i5.ClientChannel>(injectableModule.clientChannel);
    gh.singleton<_i6.FirebaseAnalytics>(injectableModule.firebaseAnalytics);
    gh.singleton<_i7.FirebaseAuth>(injectableModule.firebaseAuth);
    gh.singleton<_i8.FirebaseCrashlytics>(injectableModule.firebaseCrashlytics);
    gh.singleton<_i9.FirebaseDynamicLinks>(
        injectableModule.firebaseDynamicLinks);
    gh.singleton<_i10.FirebaseFirestore>(injectableModule.firebaseFirestore);
    gh.singleton<_i11.HealthServiceClient>(
        injectableModule.healthServiceClient);
    gh.factory<_i12.LocationControllerImpl>(
        () => _i12.LocationControllerImpl());
    gh.singleton<_i13.PlaceControllerImpl>(injectableModule.placeController);
    gh.factory<_i14.PoliciesAgreementsRepository>(() =>
        _i15.PoliciesAgreementsRepositoryImpl(gh<_i10.FirebaseFirestore>()));
    gh.singleton<_i16.SharedPreferences>(injectableModule.sharedPreferences);
    gh.singleton<_i17.StreamController<_i18.PoliciesAgreements>>(
        injectableModule.agreementsStreamCtrl);
    gh.factory<_i19.UserRepository>(
        () => _i20.UserRepositoryImpl(gh<_i10.FirebaseFirestore>()));
    gh.factory<_i21.ChatControllerImpl>(() => _i21.ChatControllerImpl(
          gh<_i4.ChatServiceClient>(),
          gh<_i10.FirebaseFirestore>(),
        ));
    gh.factory<_i22.PoliciesAgreementControllerImpl>(
        () => _i22.PoliciesAgreementControllerImpl(
              gh<_i14.PoliciesAgreementsRepository>(),
              gh<_i19.UserRepository>(),
              gh<_i17.StreamController<_i18.PoliciesAgreements>>(),
            ));
    return this;
  }
}

class _$InjectableModule extends _i23.InjectableModule {}

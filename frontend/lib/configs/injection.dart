import 'package:firebase_auth/firebase_auth.dart';
import 'package:get_it/get_it.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:grpc/grpc.dart';
import 'package:injectable/injectable.dart';
import 'package:otomo/controllers/boundary/id_token.dart';
import 'package:otomo/configs/app_config.dart';
import 'package:otomo/configs/injection.config.dart';
import 'package:otomo/controllers/id_token.dart';
import 'package:otomo/grpc/generated/chat_service_service.pbgrpc.dart';
import 'package:otomo/grpc/generated/health.pbgrpc.dart';
import 'package:otomo/grpc/generated/interceptors/auth.dart';
import 'package:otomo/grpc/generated/interceptors/logging.dart';
import 'package:otomo/grpc/generated/interceptors/retry.dart';
import 'package:otomo/tools/global_state.dart';

final getIt = GetIt.instance;

@module
abstract class InjectableModule {
  final _clientChannel = ClientChannel(
    appConfig.otomoServerHost,
    port: appConfig.otomoServerPort,
    options: ChannelOptions(
      credentials: appConfig.isSecureConnectionToOtomoServer
          ? const ChannelCredentials.secure()
          : const ChannelCredentials.insecure(),
      connectTimeout: const Duration(seconds: 10),
    ),
  );

  @singleton
  ClientChannel get clientChannel => _clientChannel;

  @singleton
  HealthServiceClient get healthServiceClient =>
      HealthServiceClient(clientChannel, interceptors: _clientInterceptors);

  @singleton
  ChatServiceClient get chatServiceClient =>
      ChatServiceClient(clientChannel, interceptors: _clientInterceptors);

  static final _firebaseAuth = FirebaseAuth.instance;
  @singleton
  FirebaseAuth get firebaseAuth => _firebaseAuth;

  @singleton
  GoogleSignIn get googleSignIn => GoogleSignIn(
        scopes: ['email'],
      );

  @singleton
  GlobalState get globalState => GlobalState.instance;

  static IdTokenController get _idTokenController =>
      IdTokenControllerImpl(_firebaseAuth);

  @singleton
  IdTokenController get idTokenController => _idTokenController;

  final List<ClientInterceptor> _clientInterceptors = [
    _loggingClientInterceptor,
    _retryClientInterceptor,
    _injectAuthHeaderClientInterceptor,
  ];

  static final _loggingClientInterceptor = LoggingClientInterceptor();
  static final _retryClientInterceptor =
      RetryOnUnavailableErrorClientInterceptor(retries: 1);
  static final _injectAuthHeaderClientInterceptor =
      InjectAuthHeaderClientInterceptor(_idTokenController);
}

@InjectableInit()
void configureInjection() => getIt.init();

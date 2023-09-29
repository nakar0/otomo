import 'package:otomo/configs/injection.dart';
import 'package:otomo/controllers/auth.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'sign_in.g.dart';

@riverpod
class SignIn extends _$SignIn {
  SignIn() : super();

  final _controller = getIt<AuthControllerImpl>();

  @override
  Future<void> build() async {}

  Future<void> signInWithGoogle() async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(_controller.signInWithGoogle);
  }

  Future<void> signInWithApple() async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(_controller.signInWithApple);
  }

  Future<void> sendSignInEmailLink(String email) async {
    state = const AsyncValue.loading();
    state =
        await AsyncValue.guard(() => _controller.sendSignInEmailLink(email));
  }

  Future<void> signInWithEmailLinkIfLinkCorrect(String link) async {
    if (!_controller.isSignInWithEmailLink(link)) return;
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() => _controller.signInWithEmailLink(link));
  }
}

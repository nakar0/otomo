import 'package:firebase_auth/firebase_auth.dart' as auth;
import 'package:injectable/injectable.dart';
import 'package:otomo/models/user.dart';

@Injectable()
class AuthController {
  AuthController(this._firebaseAuth);

  final auth.FirebaseAuth _firebaseAuth;

  Stream<User?> authStateChanges() => _firebaseAuth
      .authStateChanges()
      .map((authUser) => authUser == null ? null : User(id: authUser.uid));

  Future<String?> getIdToken() async {
    return _firebaseAuth.currentUser?.getIdToken();
  }
}

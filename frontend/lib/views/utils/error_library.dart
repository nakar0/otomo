import 'dart:io';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:grpc/grpc.dart';
import 'package:otomo/controllers/auth.dart';
import 'package:otomo/entities/app_exception.dart';

abstract class ErrorLibrary {
  static String fromAny(Object e) {
    if (e is AppException) return fromAppException(e);
    if (e is SocketException) return fromSocketException(e);
    if (e is GrpcError) return fromGrpcError(e);
    if (e is FirebaseAuthException) return fromFirebaseAuthException(e);
    return unknown;
  }

  static String fromAppException(AppException e) {
    if (e.causeIs(Cause.unknown)) return unknown;
    if (e.causeIs(Cause.networkError)) return network;
    if (e.causeIs(Cause.requiresRecentLogin)) return requiresRecentLogin;
    return unknown;
  }

  static String fromSocketException(SocketException e) {
    return network;
  }

  static String fromGrpcError(GrpcError e) {
    final details = e.details;
    final detail = details?.first;

    switch (e.code) {
      case StatusCode.invalidArgument:
        return invalidArgument;
      case StatusCode.notFound:
        return notFound;
      case StatusCode.alreadyExists:
        return alreadyExist;
      case StatusCode.permissionDenied:
        return permissionDenied;
      case StatusCode.unauthenticated:
        return unauthenticated;
      case StatusCode.resourceExhausted:
        if (detail == null) return unknown;
        if (detail is ErrorInfo) {
          if (detail.domain == Domain.message.name) return messageSendLimit;
        }
        return network;
      default:
        return network;
    }
  }

  static String fromFirebaseAuthException(FirebaseAuthException e) {
    switch (e.code) {
      case FirebaseAuthExceptionCode.requiresRecentLogin:
        return requiresRecentLogin;
      case FirebaseAuthExceptionCode.internalError:
      case FirebaseAuthExceptionCode.networkRequestFailed:
        return network;
      default:
        return unknown;
    }
  }

  // common
  static const String unknown = 'エラーが発生しました。';
  static const String network = 'ネットワークエラーが発生しました。';
  static const String invalidArgument = '入力が正しくありません。';
  static const String notFound = '該当するデータがありません。';
  static const String alreadyExist = 'すでにデータが存在します。';
  static const String permissionDenied = 'アクセス権限がありません。';
  static const String unauthenticated = '認証に失敗しました。';

  // app
  static const String requiresRecentLogin = '再ログインが必要です。';
  static const String messageLocationAnalysis = '地名の解析に失敗しました。';
  static const String messageSendLimit = 'メッセージの送信回数が上限に達しました。';
}

import 'package:dio/dio.dart';
import 'package:otomo/entities/app_exception.dart';
import 'package:otomo/entities/place_details.dart';
import 'package:otomo/tools/logger.dart';

class PlaceControllerImpl {
  PlaceControllerImpl({
    required String apiKey,
    String? androidPackageName,
    String? androidCertFingerprintSha1,
    String? iosBundleId,
  }) : _client = Dio(
          BaseOptions(
            baseUrl: 'https://maps.googleapis.com',
            queryParameters: {'key': apiKey},
            headers: {
              'x-android_package_name': androidPackageName ?? '',
              'x-android_cert_fingerprint': androidCertFingerprintSha1 ?? '',
              'x-ios-bundle-identifier': iosBundleId ?? '',
            },
          ),
        );

  final Dio _client;

  static const _detailsUrl = '/maps/api/place/details/json';

  Future<PlaceDetails> getPlaceDetails(String googlePlaceId) async {
    logger.debug('getPlaceDetails. googlePlaceId: $googlePlaceId');
    final response = await _client.get(
      _detailsUrl,
      queryParameters: {
        'place_id': googlePlaceId,
        'language': 'ja',
      },
    );

    final errorMessage = response.data['error_message'];
    if (errorMessage != null) {
      throw AppException.unknown(errorMessage);
    }

    return PlaceDetails.fromJson(response.data['result']);
  }
}

import 'package:dio/dio.dart';
import 'package:otomo/domains/entities/app_exception.dart';
import 'package:otomo/domains/entities/place_details.dart';
import 'package:otomo/tools/logger.dart';

class PlaceControllerImpl {
  PlaceControllerImpl({
    required String apiKey,
    Map<String, String> headers = const {},
  }) : _client = Dio(
          BaseOptions(
            baseUrl: 'https://maps.googleapis.com',
            queryParameters: {'key': apiKey},
            headers: headers,
          ),
        );

  PlaceControllerImpl.client(this._client);

  final Dio _client;

  static const _detailsUrl = '/maps/api/place/details/json';

  Future<PlaceDetails> getPlaceDetails(
    String googlePlaceId, {
    String language = 'en',
  }) async {
    logger.debug('getPlaceDetails. googlePlaceId: $googlePlaceId');
    final response = await _client.get(
      _detailsUrl,
      queryParameters: {
        'place_id': googlePlaceId,
        'language': language,
      },
    );

    final errorMessage = response.data['error_message'];
    if (errorMessage != null) {
      throw AppException.unknown(errorMessage);
    }

    return PlaceDetails.fromJson(response.data['result']);
  }
}

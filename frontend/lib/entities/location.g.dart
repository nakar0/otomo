// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'location.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Location _$$_LocationFromJson(Map<String, dynamic> json) => _$_Location(
      googlePlaceId: json['google_place_id'] as String,
      address: json['address'] as String,
      types: (json['types'] as List<dynamic>).map((e) => e as String).toList(),
      geometry: Geometry.fromJson(json['geometry'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$_LocationToJson(_$_Location instance) =>
    <String, dynamic>{
      'google_place_id': instance.googlePlaceId,
      'address': instance.address,
      'types': instance.types,
      'geometry': instance.geometry,
    };

_$_Geometry _$$_GeometryFromJson(Map<String, dynamic> json) => _$_Geometry(
      latLng: AppLatLng.fromJson(json['lat_lng'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$_GeometryToJson(_$_Geometry instance) =>
    <String, dynamic>{
      'lat_lng': instance.latLng,
    };

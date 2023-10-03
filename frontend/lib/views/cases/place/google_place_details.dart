import 'package:flutter/material.dart';
import 'package:otomo/configs/app_config.dart';
import 'package:otomo/entities/place_details.dart' as place_details;
import 'package:otomo/tools/logger.dart';
import 'package:otomo/views/bases/spaces/spaces.dart';
import 'package:otomo/views/bases/texts/texts.dart';
import 'package:otomo/views/cases/place/place_review_card.dart';
import 'package:otomo/views/cases/place/place_type_chips.dart';

class GooglePlaceDetails extends StatelessWidget {
  const GooglePlaceDetails({
    super.key,
    required this.place,
  });

  final place_details.PlaceDetails place;

  String _getImageUrl({
    required String photoReference,
    required int maxHeight,
  }) =>
      'https://maps.googleapis.com/maps/api/place/photo?maxheight=$maxHeight&photo_reference=$photoReference&key=${appConfig.googleMapApiKey}';

  Widget _buildWidthSized({
    required BuildContext context,
    required Widget child,
  }) {
    final size = MediaQuery.of(context).size;
    return SizedBox(
      width: size.width,
      child: child,
    );
  }

  Widget _buildBody(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        if (place.types?.isNotEmpty == true)
          _buildWidthSized(
            context: context,
            child: Wrap(
              spacing: 8,
              children: place.types!
                  .map((type) => PlaceTypeChips(type: type))
                  .toList(),
            ),
          ),
        // hours | rating
        Row(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // hours
            if (place.openingHours?.openNow != null)
              Column(
                children: [
                  const TitleSmall('Hours'),
                  Spaces.h8,
                  BodyMedium(
                    place.openingHours!.openNow! ? 'Open' : 'Closed',
                  )
                ],
              ),
            // rating
            if (place.rating != null)
              Column(
                children: [
                  const TitleSmall('Rating'),
                  Spaces.h8,
                  BodyMedium(place.rating.toString()),
                ],
              ),
          ],
        ),
        // images page view
        if (place.photos != null)
          SizedBox(
            height: 240,
            child: _buildWidthSized(
              context: context,
              child: PageView.builder(
                itemCount: place.photos!.length,
                itemBuilder: (context, index) {
                  final photo = place.photos![index];
                  return Padding(
                    padding: EdgeInsets.symmetric(horizontal: 16),
                    child: Image.network(
                      _getImageUrl(
                        photoReference: photo.photoReference,
                        maxHeight: 980,
                      ),
                      fit: BoxFit.cover,
                    ),
                  );
                },
              ),
            ),
          ),
        // overview
        if (place.editorialSummary?.overview != null)
          _buildWidthSized(
            context: context,
            child: BodyMedium(place.editorialSummary!.overview!),
          ),
        // Ratings and & Reviews
        if (place.reviews != null)
          SizedBox(
            height: 200,
            child: _buildWidthSized(
              context: context,
              child: PageView.builder(
                itemCount: place.reviews!.length,
                itemBuilder: (context, index) => PlaceReviewCard(
                  review: place.reviews![index],
                ),
              ),
            ),
          ),

        const TitleMedium('Details'),
        // phone number
        _buildWidthSized(
          context: context,
          child: Column(
            children: ListTile.divideTiles(context: context, tiles: [
              if (place.formattedPhoneNumber != null)
                ListTile(
                  leading: Icon(Icons.phone),
                  title: BodyMedium(place.formattedPhoneNumber!),
                ),
              if (place.website != null)
                ListTile(
                  leading: Icon(Icons.web),
                  title: BodyMedium(place.website!),
                ),
              if (place.formattedAddress != null)
                ListTile(
                  leading: Icon(Icons.location_on),
                  title: BodyMedium(place.formattedAddress!),
                ),
              if (place.currentOpeningHours != null)
                ExpansionTile(
                  leading: Icon(Icons.access_time),
                  title: BodyMedium('Opening Hours'),
                  children: [
                    for (var period in place.currentOpeningHours!.periods!)
                      ListTile(
                        title: BodyMedium(period.open.time),
                        subtitle: BodyMedium(period.open.day.toString()),
                      ),
                  ],
                ),
            ]).toList(),
          ),
        )
        // hours details
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    logger.debug(place);
    final theme = Theme.of(context);
    return Container(
      color: theme.colorScheme.background,
      child: Column(
        children: [
          TitleLarge(place.name ?? ''),
          _buildBody(context),
        ],
      ),
    );
  }
}

import 'package:flutter/material.dart';
import 'package:otomo/entities/place_details.dart';
import 'package:otomo/views/bases/app_bars/app_bar_title.dart';
import 'package:otomo/views/bases/layouts/edge_layout.dart';
import 'package:otomo/views/bases/sheets/sheet_form.dart';
import 'package:otomo/views/cases/place/place_review_list.dart';

class PlaceReviewsPage extends StatelessWidget {
  const PlaceReviewsPage({
    super.key,
    required this.reviews,
  });

  final List<PlaceDetailsReview> reviews;

  static Future<void> showBottomSheet({
    required BuildContext context,
    required List<PlaceDetailsReview> reviews,
  }) {
    return showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      backgroundColor: Colors.transparent,
      useSafeArea: true,
      builder: (context) => Scaffold(
        backgroundColor: Colors.transparent,
        body: SheetForm(
          shadow: false,
          child: CustomScrollView(
            slivers: slivers(context: context, reviews: reviews),
          ),
        ),
      ),
    );
  }

  static List<Widget> slivers({
    required BuildContext context,
    required List<PlaceDetailsReview> reviews,
  }) {
    final tiles = ListTile.divideTiles(
      context: context,
      tiles: [
        for (final review in reviews)
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 16),
            child: PlaceReviewListTile(review: review),
          ),
      ],
    ).toList();
    return [
      SliverAppBar(
        automaticallyImplyLeading: false,
        pinned: true,
        title: const AppBarTitle(title: 'Reviews'),
        actions: [
          IconButton(
            icon: const Icon(Icons.close_rounded),
            onPressed: () => Navigator.of(context).pop(),
          ),
        ],
      ),
      EdgeLayout(
        top: 0.0,
        sliver: true,
        child: SliverList(
          delegate: SliverChildBuilderDelegate(
            (context, index) => tiles[index],
            childCount: reviews.length,
          ),
        ),
      ),
    ];
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: CustomScrollView(
        slivers: slivers(context: context, reviews: reviews),
      ),
    );
  }
}

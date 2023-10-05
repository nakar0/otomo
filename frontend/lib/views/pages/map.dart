import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:otomo/entities/lat_lng.dart';
import 'package:otomo/entities/message.dart';
import 'package:otomo/view_models/boundary/chat.dart';
import 'package:otomo/view_models/map.dart';
import 'package:otomo/views/cases/map/app_map.dart';
import 'package:otomo/views/utils/marker_maker.dart';

class MapPage extends StatefulHookConsumerWidget {
  const MapPage({super.key});

  @override
  ConsumerState<MapPage> createState() => _MapState();
}

class _MapState extends ConsumerState<MapPage> {
  MapController? _mapController;
  Set<Marker> _markers = {};

  bool get _canUseMapController => _mapController != null;

  static const _initialCameraPosition = CameraPosition(
    target: LatLng(37.42796133580664, -122.085749655962),
    zoom: 14.4746,
  );

  Future<void> _onMapCreated(MapController controller) async {
    _mapController = controller;
    await _goCurrentLocation(ref.read(mapProvider.notifier));
  }

  Future<void> _goCurrentLocation(Map notifier) async {
    if (!_canUseMapController) return;
    final position = await notifier.getCurrentPosition();
    _mapController!.moveWithLatLng(latLng: position.latLng, zoom: 14);
  }

  Future<void> _addMarker(AnalyzedLocation loc, {bool notify = true}) async {
    final notifier = ref.read(mapProvider.notifier);
    _markers.add(await MarkerMaker.fromAnalyzedLocationWithLabel(
      context: context,
      loc: loc,
      onTap: () => notifier.focusPlace(loc.location),
    ));
    if (notify) setState(() {});
  }

  void _setMarkers(List<AnalyzedLocation> locs) async {
    _markers = {};
    for (final loc in locs) {
      await _addMarker(loc, notify: false);
    }
    setState(() {});
  }

  void _onLocationFocused(AnalyzedLocation loc) {
    if (!_canUseMapController) return;
    _addMarker(loc);
    _mapController!.moveWithLatLng(
      latLng: loc.location.geometry.latLng,
      zoom: 15,
    );
  }

  void _onTextMsgActivated(TextMessageData textMsg) {
    if (!_canUseMapController) return;

    final latLngList = AppLatLngList(
      textMsg.locationAnalysis.locations
          .map((e) => e.location.geometry.latLng)
          .toList(),
    );
    final region = latLngList.edge();
    if (region == null) return;
    _mapController!.moveWithRegion(region: region, padding: 40);
  }

  @override
  @override
  Widget build(BuildContext context) {
    ref.listen(
      mapProvider.select((value) => value.activeAnalyzedLocations),
      (prev, next) => _setMarkers(next),
    );

    final notifier = ref.read(mapProvider.notifier);

    useEffect(() {
      final focusedAnalyzedLocationStreamSub =
          notifier.focusedLocationStream.listen(_onLocationFocused);
      final activatedTextMessageStreamSub =
          notifier.activatedTextMessageStream.listen(_onTextMsgActivated);

      return () {
        focusedAnalyzedLocationStreamSub.cancel();
        activatedTextMessageStreamSub.cancel();
      };
    }, const []);

    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: AppMap(
        initialCameraPosition: _initialCameraPosition,
        onMapCreated: _onMapCreated,
        markers: _markers,
      ),
      floatingActionButton: FloatingActionButton(
        mini: true,
        onPressed: () => _goCurrentLocation(notifier),
        child: const Icon(Icons.location_searching_rounded),
      ),
      floatingActionButtonLocation: FloatingActionButtonLocation.endTop,
    );
  }
}
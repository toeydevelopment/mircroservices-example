import 'package:flutter/material.dart';
import 'package:mobile/core/constants.dart';

class BaseAppbar extends StatelessWidget implements PreferredSizeWidget {
  final String title;
  const BaseAppbar({Key? key, required this.title}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AppBar(
      title: Text(
        title,
        style: const TextStyle(color: Colors.white),
      ),
      centerTitle: true,
      backgroundColor: kPrimaryColor,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}

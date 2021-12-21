import 'package:flutter/material.dart';

class PartyCardWidget extends StatelessWidget {
  final String name;
  final String desc;
  final String seats;
  final bool isJoined;

  const PartyCardWidget({
    Key? key,
    required this.name,
    required this.desc,
    required this.seats,
    required this.isJoined,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 10,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
      child: Column(
        children: [
          ClipRRect(
            borderRadius: const BorderRadius.only(
              topLeft: Radius.circular(8),
              topRight: Radius.circular(8),
            ),
            child: Image.network("https://via.placeholder.com/150"),
          ),
          const SizedBox(
            height: 10,
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 10),
            child: Column(
              children: [
                Text(name),
                const SizedBox(
                  height: 5,
                ),
                Text(desc),
                const SizedBox(
                  height: 10,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(seats),
                    if (isJoined)
                      ElevatedButton(
                        onPressed: () {},
                        child: const Text("Cancel"),
                      )
                    else
                      ElevatedButton(
                        onPressed: () {},
                        child: const Text("Join"),
                      )
                  ],
                ),
              ],
            ),
          ),
          const SizedBox(
            height: 10,
          )
        ],
      ),
    );
  }
}

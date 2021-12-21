import 'package:flutter/material.dart';

class CreatePartyPage extends StatelessWidget {
  const CreatePartyPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Form(
              child: Column(
            children: [
              TextFormField(
                decoration: const InputDecoration(labelText: "Name"),
              ),
              const SizedBox(
                height: 10,
              ),
              TextFormField(
                decoration: const InputDecoration(labelText: "Description"),
              ),
              const SizedBox(
                height: 10,
              ),
              TextFormField(
                keyboardType: TextInputType.number,
                decoration: const InputDecoration(
                  labelText: "Seats",
                ),
              ),
            ],
          ))
        ],
      ),
    );
  }
}

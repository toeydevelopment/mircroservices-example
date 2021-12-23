import React, { useState } from "react";
import { TextField, Button } from ".";
import useAuth from "../hooks/useAuth";
const SignUpForm = () => {
  const { signup } = useAuth();
  const [form, setForm] = useState({ email: "", password: "" });

  return (
    <div className="border border-2 border-grey-300 m-10 shadow-xl w-5/6 max-w-xl md:w-4/6 md:h-2/6 p-6 rounded-md flex flex-col justify-center items-center">
      <h2 className="text-xl text-center mb-4">Sign Up</h2>
      <TextField
        label="Email"
        placeholder="...@gmail.com"
        type="email"
        onChanged={(e) => {
          setForm((p) => ({ ...p, email: e }));
        }}
      />
      <TextField
        label="Password"
        placeholder=""
        type="password"
        onChanged={(e) => {
          setForm((p) => ({ ...p, password: e }));
        }}
      />
      <TextField
        label="Re-password"
        placeholder=""
        type="password"
        onChanged={(e) => {}}
      />
      <Button onClicked={() => signup(form.email, form.password)}>
        {" "}
        Sign Up{" "}
      </Button>
    </div>
  );
};

export default SignUpForm;

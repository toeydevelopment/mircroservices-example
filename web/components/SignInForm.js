import { useRouter } from "next/router";
import axios from "axios";
import React, { useCallback, useState } from "react";
import { TextField, Button } from ".";
import useAuth from "../hooks/useAuth";

const SignInForm = () => {
  const router = useRouter();
  const { login } = useAuth();

  const [form, setForm] = useState({
    email: "",
    password: "",
  });

  const submit = useCallback(() => {
    login(form.email, form.password);
  }, [form]);

  return (
    <div className="border border-2 border-grey-300 m-10 shadow-xl w-5/6 max-w-xl md:w-4/6 md:h-2/6 p-6 rounded-md flex flex-col justify-center items-center">
      <h2 className="text-xl text-center mb-4">Sign In</h2>
      <TextField
        label="Email"
        placeholder="...@gmail.com"
        type="email"
        onChanged={(e) => {
          setForm((p) => {
            return {
              ...p,
              email: e,
            };
          });
        }}
      />
      <TextField
        label="Password"
        placeholder=""
        type="password"
        onChanged={(e) => {
          setForm((p) => {
            return {
              ...p,
              password: e,
            };
          });
        }}
      />
      <div className="mb-2 w-1/2">
        <Button onClicked={submit}> Sign In </Button>
      </div>
      <div className="w-1/2">
        <Button
          onClicked={() => {
            router.push("/sign-up");
          }}
        >
          {" "}
          Sign Up{" "}
        </Button>
      </div>
    </div>
  );
};

export default SignInForm;

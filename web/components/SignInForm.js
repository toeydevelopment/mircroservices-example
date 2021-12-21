import { useRouter } from "next/router";
import React from "react";
import { TextField, Button } from ".";

const SignInForm = () => {
  const router = useRouter();

  return (
    <div className="border border-2 border-grey-300 m-10 shadow-xl w-5/6 max-w-xl md:w-4/6 md:h-2/6 p-6 rounded-md flex flex-col justify-center items-center">
      <h2 className="text-xl text-center mb-4">Sign Up</h2>
      <TextField label="Email" placeholder="...@gmail.com" type="email" />
      <TextField label="Password" placeholder="" type="password" />
      <div className="mb-2 w-1/2">
        <Button> Sign In </Button>
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

import React from "react";
import { TextField, Button } from ".";
const SignUpForm = () => {
  return (
    <div className="border border-2 border-grey-300 m-10 shadow-xl w-5/6 max-w-xl md:w-4/6 md:h-2/6 p-6 rounded-md flex flex-col justify-center items-center">
      <h2 className="text-xl text-center mb-4">Sign Up</h2>
      <TextField label="Email" placeholder="...@gmail.com" type="email" />
      <TextField label="Password" placeholder="" type="password" />
      <TextField label="Re-password" placeholder="" type="password" />
      <Button> Sign Up </Button>
    </div>
  );
};

export default SignUpForm;

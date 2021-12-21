import React from "react";

const Button = ({ children, onClicked }) => {
  return (
    <button
      onClick={onClicked}
      className="hover:border hover:border-2 hover:border-primary bg-primary hover:bg-yellow-300 text-white hover:text-primary font-bold py-2 px-4 rounded min-w-full"
    >
      {children}
    </button>
  );
};

export default Button;

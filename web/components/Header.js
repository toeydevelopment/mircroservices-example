import { useRouter } from "next/router";
import React, { useCallback } from "react";

const Header = ({ children }) => {
  const router = useRouter();

  return (
    <header className="relative w-full py-5 bg-primary uppercase text-white text-xl md:text-4xl xl:text-5xl text-center shadow-md tracking-wider">
      {router.pathname !== "/" && (
        <div
          onClick={() => {
            router.back();
          }}
          className="inline-block absolute left-5 w-10 h-10 hover:bg-gray-200 rounded-full cursor-pointer flex justify-center items-center"
        >
          &#8249;
        </div>
      )}{" "}
      {children}
    </header>
  );
};

export default Header;

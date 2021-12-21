import { useEffect, useState } from "react";

/**
 * useAuth is simple authentication hook
 * @returns isAuth indicated wheater user logged in or not
 */
const useAuth = () => {
  const [auth, setAuth] = useState("");

  useEffect(() => {
    setAuth(localStorage.getItem("auth") ?? "");
    console.log(auth);
  }, [auth]);

  return {
    isAuth: auth !== "",
    setAuth,
  };
};

export default useAuth;

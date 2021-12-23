import axios from "axios";
import jwtDecode from "jwt-decode";
import { useRouter } from "next/router";
import { useEffect, useMemo, useState } from "react";

/**
 * useAuth is simple authentication hook
 * @returns isAuth indicated wheater user logged in or not
 */
const useAuth = () => {
  const [auth, setAuth] = useState("");
  const [email, setEmail] = useState("");
  const router = useRouter();

  useEffect(() => {
    setAuth(localStorage.getItem("auth") ?? "");
    setEmail(localStorage.getItem("email") ?? "");
  }, [auth, email]);

  const login = async (email, password) => {
    try {
      const res = await axios.post(
        "/api/authentication-command-service/auth/signin/",
        {
          email,
          password,
        }
      );

      const { email: nemail } = jwtDecode(res.data["access_token"]);

      setEmail(nemail);

      localStorage.setItem("email", nemail);
      localStorage.setItem("auth", res.data["access_token"]);
      setAuth(res.data["access_token"]);
      router.reload();
    } catch (error) {
      alert(error);
    }
  };

  const signup = async (email, password) => {
    try {
      const res = await axios.post(
        "/api/authentication-command-service/auth/signup/",
        {
          email,
          password,
        }
      );

      const { email: nemail } = jwtDecode(res.data["access_token"]);

      setEmail(nemail);

      localStorage.setItem("email", nemail);
      localStorage.setItem("auth", res.data["access_token"]);

      router.replace("/");
      setAuth(res.data["access_token"]);
    } catch (error) {
      alert(error);
    }
  };

  return useMemo(
    () => ({
      isAuth: auth !== "",
      setAuth,
      login,
      email,
      signup,
    }),
    [auth, email]
  );
};

export default useAuth;

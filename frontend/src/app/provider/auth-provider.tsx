"use client";

import { useRouter } from "next/navigation";
import React, { useEffect } from "react";
// import { useJwt } from "react-jwt";

const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const token =
    (typeof window != "undefined" && localStorage.getItem("token")) || "";
  //   const { isExpired } = useJwt(token!);
  const router = useRouter(); // useNavigate replaces useHistory in React Router v6

  useEffect(() => {
    if (!token) {
      router.push("/login"); // Redirect to login page if no token or token is expired
    } else {
      router.push("/");
    }
    console.log(token);
  }, [token]);

  return children;
};

export default AuthProvider;

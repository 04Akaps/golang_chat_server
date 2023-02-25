import React, { useState, useEffect } from "react";
import axios from "axios";
import styled from "styled-components";
import { useNavigate, Link } from "react-router-dom";
import Logo from "../assets/logo.svg";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { loginRoute } from "../utils/APIRoutes";
import { isExistCookie } from "../utils/CookieChecker";

import { useCookies } from "react-cookie";

export default function Login() {
  const [cookies, setCookie] = useCookies(["auth"]);
  const navigate = useNavigate();
  const toastOptions = {
    position: "bottom-right",
    autoClose: 8000,
    pauseOnHover: true,
    draggable: true,
    theme: "dark",
  };
  useEffect(() => {
    if (isExistCookie(cookies)) {
      navigate("/");
    }
  }, []);

  const handleOauth2 = async (oAuth) => {
    return window.location.replace(`http://localhost:8080/auth/login/${oAuth}`);
  };

  return (
    <>
      <FormContainer>
        <div>
          <div className="brand">
            <img src={Logo} alt="logo" />
            <h1>Free Chat</h1>
          </div>
          <button
            type="button"
            onClick={() => {
              handleOauth2("google");
            }}
          >
            Google
          </button>
          <button
            type="button"
            onClick={() => {
              handleOauth2("facebook");
            }}
          >
            FaceBook
          </button>
          <button
            type="button"
            onClick={() => {
              handleOauth2("github");
            }}
          >
            Github
          </button>

          {/* <button type="submit">Log In</button> */}
        </div>
      </FormContainer>
      <ToastContainer />
    </>
  );
}

const FormContainer = styled.div`
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;
  align-items: center;
  background-color: #131324;
  .brand {
    display: flex;
    align-items: center;
    gap: 1rem;
    justify-content: center;
    img {
      height: 5rem;
    }
    h1 {
      color: white;
      text-transform: uppercase;
    }
  }

  div {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    background-color: #00000076;
    border-radius: 2rem;
    padding: 5rem;
  }
  input {
    background-color: transparent;
    padding: 1rem;
    border: 0.1rem solid #4e0eff;
    border-radius: 0.4rem;
    color: white;
    width: 100%;
    font-size: 1rem;
    &:focus {
      border: 0.1rem solid #997af0;
      outline: none;
    }
  }
  button {
    background-color: #4e0eff;
    color: white;
    padding: 1rem 2rem;
    border: none;
    font-weight: bold;
    cursor: pointer;
    border-radius: 0.4rem;
    font-size: 1rem;
    text-transform: uppercase;
    &:hover {
      background-color: #4e0eff;
    }
  }
  span {
    color: white;
    text-transform: uppercase;
    a {
      color: #4e0eff;
      text-decoration: none;
      font-weight: bold;
    }
  }
`;

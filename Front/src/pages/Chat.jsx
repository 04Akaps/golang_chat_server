import React, { useEffect, useState, useRef, useContext } from "react";

import { useNavigate } from "react-router-dom";

import styled from "styled-components";
import { socketHost } from "../utils/APIRoutes";

import { useCookies } from "react-cookie";
import { isExistCookie } from "../utils/CookieChecker";
import { Buffer } from "buffer";
import ChatContainer from "../components/ChatContainer";
import Contacts from "../components/Contacts";
import Welcome from "../components/Welcome";
import { useGlobalData } from "../context/context";

export default function Chat() {
  const [cookies] = useCookies(["auth"]);

  const storage = useGlobalData();
  const navigate = useNavigate();

  const [initLoading, setInitLoading] = useState(false);

  useEffect(() => {
    if (!isExistCookie(cookies)) {
      navigate("/login");
    } else {
      const socket = new WebSocket(socketHost);

      storage.setGlobalData(
        socket,
        JSON.parse(Buffer.from(cookies.auth, "base64").toString("utf8")).name
      );

      socket.onmessage = function (e) {
        console.log(e);
      };

      setInitLoading(true);
    }
  }, []);

  return (
    <>
      <Container>
        <div className="container">
          <Contacts userName={storage.userName} />
          {!initLoading ? (
            <Welcome />
          ) : (
            <ChatContainer userName={storage.userName} />
          )}
        </div>
        {/* <span>Your Name : {userName}</span>
        <input
          value={inputValue}
          onChange={(e) => {
            setInputValue(e.target.value);
          }}
          style={{
            background: "#fff",
            width: "300px",
            padding: "20px",
          }}
        ></input>
        <button
          onClick={() => {
            socketWeb.send(inputValue);
            setInputValue("");
          }}
        >
          Submit
        </button> */}
        {/* <div className="container">
          <Contacts contacts={contacts} changeChat={handleChatChange} />
          {currentChat === undefined ? (
            <Welcome />
          ) : (
            <ChatContainer currentChat={currentChat} socket={socket} />
          )}
        </div> */}
      </Container>
    </>
  );
}

const Container = styled.div`
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;
  align-items: center;
  background-color: #131324;
  .container {
    height: 85vh;
    width: 85vw;
    background-color: #00000076;
    display: grid;
    grid-template-columns: 25% 75%;
    @media screen and (min-width: 720px) and (max-width: 1080px) {
      grid-template-columns: 35% 65%;
    }
  }
`;

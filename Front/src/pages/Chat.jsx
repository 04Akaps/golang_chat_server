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
  const [cookies, removeToken] = useCookies(["auth"]);

  const storage = useGlobalData();
  const navigate = useNavigate();

  const [initLoading, setInitLoading] = useState(false);
  const [chatContents, setChatContents] = useState([]);
  const [profileImg, setProfileImg] = useState("");

  useEffect(() => {
    if (!isExistCookie(cookies)) {
      navigate("/login");
    } else {
      const socket = new WebSocket(socketHost);

      const authData = JSON.parse(
        Buffer.from(cookies.auth, "base64").toString("utf8")
      );
      storage.setGlobalData(socket, authData.name);

      setProfileImg(authData.avatar_url);

      setInitLoading(true);
    }
  }, []);

  if (storage.socket) {
    storage.socket.onmessage = function (e) {
      const receiveData = JSON.parse(e.data);
      if (chatContents.length === 0) {
        setChatContents([receiveData]);
      } else {
        setChatContents([...chatContents, receiveData]);
      }
    };

    storage.socket.onclose = function (e) {
      alert("서버가 닫혀있기 떄문에 로그아웃 됩니다.");

      document.cookie =
        "auth" +
        "=" +
        ("/" ? ";path=" + "/" : "") +
        ";expires=Thu, 01 Jan 1970 00:00:01 GMT";

      window.location.replace("/login");
    };
  }

  return (
    <>
      <Container>
        <div className="container">
          <Contacts profileImg={profileImg} userName={storage.userName} />
          {!initLoading && !storage.room ? (
            <Welcome userName={storage.userName} />
          ) : (
            <ChatContainer
              userName={storage.userName}
              socket={storage.socket}
              chatContents={chatContents}
              profileImg={profileImg}
            />
          )}
        </div>
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

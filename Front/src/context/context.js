import React, { useContext } from "react";

class GlobalData {
  socket;
  userName;
  roomNumber;

  setGlobalData = (socket, userName) => {
    this.socket = socket;
    this.userName = userName;
  };

  setRoom = (room) => {
    this.roomNumber = room;
  };
}

export const GlobalDataContext = React.createContext(new GlobalData());

export const useGlobalData = () => {
  return useContext(GlobalDataContext);
};

export default GlobalData;

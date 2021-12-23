import axios from "axios";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";

const useParty = () => {
  const [party, setParty] = useState([]);
  const router = useRouter();
  let cursor = "";

  useEffect(async () => {
    try {
      cursor = "";
      const res = await axios.get("/api/party-query-service/?cursor=" + cursor);

      setParty(res.data["data"]);
      cursor = res.data["cursor"];
    } catch (error) {
      alert(error);
    }
  }, []);

  const loadmore = async () => {
    const res = await axios.get("/api/party-query-service/?cursor=" + cursor);

    setParty(...party, res.data["data"]);
    cursor = res.data["cursor"];
  };

  const joinParty = async (id) => {
    try {
      const token = localStorage.getItem("auth");
      const email = localStorage.getItem("email");

      await axios.post(
        `/api/party-orchestration-service/${id}/join`,
        {},
        {
          headers: {
            Authorization: "Bearer " + token,
          },
        }
      );

      setParty((prev) =>
        prev.map((p) =>
          p.id === id
            ? {
                ...p,
                seat: p.seat++,
                joined: [...p.joined, email],
              }
            : p
        )
      );
    } catch (error) {
      alert(error);
    }
  };

  const unjoinParty = async (id) => {
    const token = localStorage.getItem("auth");
    const email = localStorage.getItem("email");

    try {
      await axios.post(
        `/api/party-orchestration-service/${id}/unjoin`,
        {},
        {
          headers: {
            Authorization: "Bearer " + token,
          },
        }
      );
      setParty((prev) =>
        prev.map((p) =>
          p.id === id
            ? {
                ...p,
                seat: p.seat--,
                joined: p.joined.filter((_) => _ != email),
              }
            : p
        )
      );
    } catch (error) {
      alert(error);
    }
  };

  const createParty = async (name, desc, seat) => {
    const token = localStorage.getItem("auth");

    try {
      await axios.post(
        `/api/party-orchestration-service/`,
        {
          name,
          description: desc,
          seat_limit: seat,
        },
        {
          headers: {
            Authorization: "Bearer " + token,
          },
        }
      );

      router.replace("/");
    } catch (error) {
      alert(error);
    }
  };

  return { loadmore, party, joinParty, unjoinParty, createParty };
};

export default useParty;

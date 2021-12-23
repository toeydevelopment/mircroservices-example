import React from "react";
import { Button } from ".";

const PartyCard = ({
  imgURL,
  name,
  desc,
  isJoined,
  onJoin,
  onCancel,
  avaliable,
}) => {
  return (
    <div className="max-w-xl max-h-xl rounded overflow-hidden shadow-lg">
      <div
        className="h-56 w-full overflow-hidden bg-cover bg-center bg-no-repeat"
        style={{
          backgroundImage: `url("${
            imgURL ?? "https://via.placeholder.com/150"
          }")`,
        }}
      ></div>
      <div className="px-6 py-4">
        <div className="font-bold text-xl mb-2">
          {name ?? "The Coldest Sunset"}
        </div>
        <p className="text-gray-700 text-base">
          {desc ??
            `Lorem ipsum dolor sit amet, consectetur adipisicing elit. Voluptatibus
          quia, nulla! Maiores et perferendis eaque, exercitationem praesentium
          nihil.`}
        </p>
      </div>
      <div className="px-6 pt-4 pb-2 w-full flex justify-between">
        <div> {avaliable ?? "0 / 6"}</div>
        <div className="w-1/3">
          {!isJoined ? (
            <Button onClicked={onJoin}>Join</Button>
          ) : (
            <Button onClicked={onCancel}>Cancel</Button>
          )}
        </div>
      </div>
    </div>
  );
};

export default PartyCard;

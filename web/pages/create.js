import React, { useState } from "react";
import { Button, Header, TextField } from "../components";
import useParty from "../hooks/useParty";

const CreatePage = () => {
  const [form, setForm] = useState({
    name: "",
    desc: "",
    seat: 0,
  });

  const { createParty } = useParty();

  return (
    <>
      <Header>Create Party</Header>
      <div className="w-5/6 md:w-2/6 mx-auto mt-10">
        <div className="mb-2">
          <label className="block uppercase tracking-wide text-gray-700 text-sm font-bold mb-2">
            File Upload
          </label>
          <div className="flex items-center justify-center w-full">
            <label className="flex flex-col w-full h-32 border-4 border-blue-200 border-dashed hover:bg-gray-100 hover:border-gray-300">
              <div className="flex flex-col items-center justify-center pt-7">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="w-8 h-8 text-gray-400 group-hover:text-gray-600"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
                  />
                </svg>
                <p className="pt-1 text-sm tracking-wider text-gray-400 group-hover:text-gray-600">
                  Attach a file
                </p>
              </div>
              <input
                type="file"
                accept="image/png, image/gif, image/jpeg"
                className="opacity-0"
              />
            </label>
          </div>
        </div>
        <TextField
          label={"Name"}
          type="text"
          onChanged={(e) => {
            setForm((p) => ({ ...p, name: e }));
          }}
        />
        <TextField
          label={"Description"}
          type="text"
          onChanged={(e) => {
            setForm((p) => ({ ...p, desc: e }));
          }}
        />
        <TextField
          label={"Seats"}
          type="number"
          onChanged={(e) => {
            setForm((p) => ({ ...p, seat: Number(e) }));
          }}
        />
        <Button onClicked={() => createParty(form.name, form.desc, form.seat)}>
          Submit
        </Button>
      </div>
    </>
  );
};

export default CreatePage;

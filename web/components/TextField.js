import React from "react";

const TextField = ({ type, label, placeholder, onChanged }) => {
  return (
    <div className="w-full">
      <label
        className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
        htmlFor="grid-first-name"
      >
        {label}
      </label>
      <input
        onChange={(e) => {
          e.preventDefault();
          onChanged(e.target.value);
        }}
        className="appearance-none block w-full bg-gray-200 text-gray-700 border rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
        id="grid-first-name"
        type={type}
        placeholder={placeholder}
      />
    </div>
  );
};

export default TextField;

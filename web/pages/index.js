import { AddButton, Header, SignInForm } from "../components";
import PartyCard from "../components/PartyCard";
import useAuth from "../hooks/useAuth";

export default function Home() {
  const { isAuth } = useAuth();

  if (isAuth) {
    return (
      <div className="w-screen h-screen flex justify-center items-center">
        <SignInForm />
      </div>
    );
  }

  return (
    <div>
      <Header>All Parties</Header>
      <div className="md:container md:mx-auto">
        <div className="w-full grid grid-cols-1 md:grid-cols-3 xl:grid-cols-4 gap-2 md:gap-3 p-6 md:p-10">
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
          <PartyCard />
        </div>
        <div className="fixed bottom-10 md:bottom-20 right-10">
          <AddButton />
        </div>
      </div>
    </div>
  );
}

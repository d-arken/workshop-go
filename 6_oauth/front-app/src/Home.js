import {useKindeAuth} from "@kinde-oss/kinde-auth-react";
import {useEffect, useMemo, useState} from "react";

export function Home() {

    const { user, isAuthenticated, isLoading, login, register, logout, getToken } = useKindeAuth();

    const [token, setToken] = useState("")
    useEffect(() => {
        if (isAuthenticated) {
          getToken().then((token) => setToken(token))
        }
    }, [isAuthenticated]);


    if (isLoading) {
        return <p>Loading</p>;
    }


    return (
        <>
            {
                isAuthenticated ?
                    <>
                        <div>
                            <h2>{user.given_name}</h2>
                            <p>{user.email}</p>
                        </div>
                        <button onClick={logout}>Logout</button>
                        <p>{token}</p>

                    </>

                    :

                    <div>
                        <button onClick={register} type="button">Sign up</button>
                        <button onClick={login} type="button">Sign In</button>
                    </div>
            }
        </>
    )
}
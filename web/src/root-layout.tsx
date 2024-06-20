import Routes from "./routes";
import {Link} from "wouter";
import {FC, useContext} from "react";
import {AuthContext} from "./api/auth-provider.tsx";
import {User} from "./api/authenticate.ts";
import useConfig from "./api/use-config.ts";


const Login: FC = () => <p>(<Link href="login">log in</Link>)</p>
const UserId: FC<{ user: User }> = ({user}) => <><p>(User: <Link to="/profile">{user.name}</Link>)</p></>

export default function RootLayout() {
    const {user} = useContext(AuthContext)
    const {title} = useConfig()

    return <>
        <h1>{ title }</h1>
        {user == null ? <Login/> : <UserId user={user}/>}
        <hr/>
        <br/>
        <Routes/>
    </>
}
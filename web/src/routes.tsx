import {Redirect, Route, Switch} from "wouter";
import App from "./pages/reading/app.tsx";
import Login from "./pages/login/login";
import Profile from "./pages/profile/profile.tsx";

const Routes = () => <>
    <Switch>
        <Route path="/">
            <Redirect to="/1"/>
        </Route>
        <Route path={/^[/](?<readingStr>[0-9]+)$/}>
            <App/>
        </Route>
        <Route path="/login">
            <Login/>
        </Route>
        <Route path="/profile">
            <Profile/>
        </Route>
    </Switch>
</>

export default Routes
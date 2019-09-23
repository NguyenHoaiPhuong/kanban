import React, { useEffect } from 'react'

import Avatar from '@material-ui/core/Avatar';
import CssBaseline from '@material-ui/core/CssBaseline';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';

import Copyright from './copyright'
import NameInput from './nameInput'
import PasswordInput from './passwordInput'
import SubmitButton from './submitButton'
import useStyles from './styles'

export default function Signin(props) {
    const classes = useStyles();

    function handleSigninSubmit(event) {
        event.preventDefault();

        let user = {
            name: event.target.username.value,
            password: event.target.password.value
        }

        // Fake authentication
        if (user.name === 'admin' && user.password === 'admin') {
            localStorage.setItem('user', user)
            props.history.replace('/')
        }
    }

    useEffect(() => {
        if (localStorage.getItem('user') != null) {
            props.history.replace('/')
        }
    }, []);

    return (
        <Grid container component="main" className={classes.root}>
            <CssBaseline />

            <Grid item xs={false} sm={4} md={7} className={classes.image} />

            <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
                <div className={classes.paper}>
                    <Avatar className={classes.avatar}>
                        <LockOutlinedIcon />
                    </Avatar>

                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>

                    <form className={classes.form} noValidate onSubmit={handleSigninSubmit}>
                        <NameInput
                            id="username"
                            name="username"
                            label="User Name or Email Address"
                            autoComplete="username"
                        />
                        <PasswordInput
                            name="password"
                            label="Password"
                            type="password"
                            id="password"
                            autoComplete={true}
                        />
                        <FormControlLabel
                            control={<Checkbox value="remember" color="primary" />}
                            label="Remember me"
                        />
                        <SubmitButton content="Sign In"/>

                        <Grid container>
                            <Grid item xs>
                                <Link href="#" variant="body2">
                                    Forgot password?
                                </Link>
                            </Grid>

                            <Grid item>
                                <Link href="/signup" variant="body2">
                                    {"Don't have an account? Sign Up"}
                                </Link>
                            </Grid>
                        </Grid>

                        <Box mt={5}>
                            <Copyright />
                        </Box>
                    </form>
                </div>
            </Grid>
        </Grid>
    );
}

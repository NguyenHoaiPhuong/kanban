import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
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

export default function Signup() {
    const classes = useStyles();
  
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
                        Sign up
                    </Typography>

                    <form className={classes.form} noValidate>
                        <Grid container spacing={2}>
                            <Grid item xs={12} sm={6}>
                                <NameInput
                                    id="firstName"
                                    name="firstName"
                                    label="First Name"
                                    autoComplete="fname"
                                />
                            </Grid>
                            <Grid item xs={12} sm={6}>
                                <NameInput
                                    id="lastName"
                                    name="lastName"
                                    label="Last Name"
                                    autoComplete="lname"
                                />
                            </Grid>
                        </Grid>
                        
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
                        <PasswordInput                            
                            name="retypepassword"
                            label="Re-type Password"
                            type="password"
                            id="retypepassword"
                            autoComplete={false}
                        />
                        <FormControlLabel
                            control={<Checkbox value="confirm" color="primary" />}
                            label="I want to receive information, market promotions and updates via email"
                        />
                        <SubmitButton content="Sign Up"/>

                        <Box mt={5}>
                            <Copyright />
                        </Box>
                    </form>
                </div>
            </Grid>
        </Grid>
    );
}
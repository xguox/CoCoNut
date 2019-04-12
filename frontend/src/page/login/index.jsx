import React, { Component } from 'react';

import './login.scss';

export default class Login extends Component {
    render() {
        return(
            <div className='admin-login--container'>
              <h1 className='admin-login--title'>COCONUT</h1>
              <form className='admin-login--form'>
                <input placeholder='name' />
                <input placeholder='pwd' />
                <button>登录</button>
              </form>
            </div>
        )
    }
}
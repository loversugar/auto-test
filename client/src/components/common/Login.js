import React, { Component } from 'react';
import {setCookie} from "../../helpers/cookies";
import '../../style/login.less';
import { Form, Icon, Input, Button, Checkbox, message, Spin } from 'antd';
import fs from 'fs';

const FormItem = Form.Item;

const client_id = 'b7f8065ab0c7188c2a21';
const authorize_uri = 'https://github.com/login/oauth/authorize';
const redirect_uri = 'http://localhost:8080/oauth/redirect';

const users = [{
    username:'admin',
    password:'admin'
},{
    username:'Totti',
    password:'Totti'
}];

function PatchUser(values) {  //匹配用户
    const {username, password} = values;
    return users.find(user => user.username === username && user.password === password);
}

class NormalLoginForm extends Component {
    state = {
        isLoding:false,
    };
    handleSubmit = (e) => {
        e.preventDefault();
        this.props.form.validateFields((err, values) => {
            // fs.mkdir('./test',(err)=>{
            //     if (err) {
            //         console.log("创建失败");
            //         return;
            //     }
            //     console.log("创建成功");
                
            // });
            // fs.writeFile('./try4.txt', "HelloWorld!");
            // // 保存行为
            // // 打开文件
            // fs.open('./try4.txt', `w`, function(err, fd) {
            //     if (err) {
            //         throw err;
            //     }
            //     console.log('open file success.');
            //     var buffer = new Buffer('shiyanlou');
            //     // 读取文件
            //     fs.write(fd, buffer, 0, 6, 0, function(err, bytesWritten, buffer) {
            //         if (err) {
            //             throw err;
            //         }

            //         console.log('write success.');
            //         // 打印出buffer中存入的数据
            //         console.log(bytesWritten, buffer.slice(0, bytesWritten).toString());

            //         // 关闭文件
            //         fs.close(fd);
            //     });
            // });

            if (!err) {
                fs.writeFile('./try4.txt', "HelloWorld");
                console.log('Received values of form: ', values);
                if(PatchUser(values)){
                    this.setState({
                        isLoding: true,
                    });

                    setCookie('mspa_user',JSON.stringify(values));
                    message.success('login successed!'); //成功信息

                    let that = this;
                    setTimeout(function() { //延迟进入
                        that.props.history.push({pathname:'/app',state:values});
                    }, 2000);

                }else{
                    message.error('login failed!'); //失败信息
                }
            }
        });
    };

    render() {
        const { getFieldDecorator } = this.props.form;
        return (
            this.state.isLoding?<Spin size="large" className="loading" />:
            <div className="login">
                <div className="login-form">
                    <div className="login-logo">
                        <div className="login-name">MSPA</div>
                    </div>
                    <Form onSubmit={this.handleSubmit} style={{maxWidth: '300px'}}>
                        <FormItem>
                            {getFieldDecorator('username', {
                                rules: [{ required: true, message: '请输入用户名!' }],
                            })(
                                <Input prefix={<Icon type="user" style={{ fontSize: 13 }} />} placeholder="用户名 (admin)" />
                            )}
                        </FormItem>
                        <FormItem>
                            {getFieldDecorator('password', {
                                rules: [{ required: true, message: '请输入密码!' }],
                            })(
                                <Input prefix={<Icon type="lock" style={{ fontSize: 13 }} />} type="password" placeholder="密码 (admin)" />
                            )}
                        </FormItem>
                        <FormItem style={{marginBottom:'0'}}>
                            {getFieldDecorator('remember', {
                                valuePropName: 'checked',
                                initialValue: true,
                            })(
                                <Checkbox>记住我</Checkbox>
                            )}
                            <a className="login-form-forgot" href="" style={{float:'right'}}>忘记密码?</a>
                            <Button type="primary" htmlType="submit" className="login-form-button" style={{width: '100%'}}>
                                登录
                            </Button>
                            Or <a href="">现在就去注册!</a>
                        </FormItem>
                    </Form>
                    <a className="githubUrl" href={`${authorize_uri}?client_id=${client_id}&redirect_uri=${redirect_uri}`}> </a>
                </div>
            </div>
        );
    }
}

const Login = Form.create()(NormalLoginForm);
export default Login;
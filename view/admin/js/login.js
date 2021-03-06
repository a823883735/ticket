new Vue({
	el: '#app',
	data: function() {
		// return { }
		var checkAge = (rule, value, callback) => {
			if (!value) {
				return callback(new Error('年龄不能为空'));
			}
			setTimeout(() => {
				if (!Number.isInteger(value)) {
					callback(new Error('请输入数字值'));
				} else {
					if (value < 18) {
						callback(new Error('必须年满18岁'));
					} else {
						callback();
					}
				}
			}, 1000);
		};
		var validatePass = (rule, value, callback) => {
			if (value === '') {
				callback(new Error('请输入密码'));
			} else {
				if (this.ruleForm.checkPass !== '') {
					this.$refs.ruleForm.validateField('checkPass');
				}
				callback();
			}
		};
		var validatePass2 = (rule, value, callback) => {
			if (value === '') {
				callback(new Error('请再次输入密码'));
			} else if (value !== this.ruleForm.pass) {
				callback(new Error('两次输入密码不一致!'));
			} else {
				callback();
			}
		};
		var validateUserName = (rule, value, callback) => {
			if (value === '') {
				callback(new Error('请输入账号'));
			} else {
				// if (this.ruleForm.userName !== '') {
				//   this.$refs.ruleForm.validateField('userName');
				// }
				callback();
			}
		}
		return {
			ruleForm: {
				userName: 'admin',
				pass: '123456',
				// checkPass: '',
				// age: ''
			},
			rules: {
				userName: [{
					validator: validateUserName,
					trigger: 'blur'
				}],
				pass: [{
					validator: validatePass,
					trigger: 'blur'
				}],
				// checkPass: [
				//   { validator: validatePass2, trigger: 'blur' }
				// ],
				// age: [
				//   { validator: checkAge, trigger: 'blur' }
				// ]
			}
		};
	},
	methods: {
		submitForm(formName) {
			var that = this;
			that.$refs[formName].validate((valid) => {
				if (valid) {
					axios.post('http://106.13.237.16:10000/v1/admin/login', Qs.stringify({
						userName: that.ruleForm.userName,
						password: that.ruleForm.pass
					}), {
						headers: {
							"Content-Type": "application/x-www-form-urlencoded",
						}
					}).then(function(res) {
						var ctx = res.data;
						if (ctx.code == 200) {
							localStorage.setItem('id', that.ruleForm.userName);
							that.$alert('登录成功', '提示', {
								confirmButtonText: '确定',
								callback: action => {
									window.location.href = "index.html";
									// that.$message({
									//   type: 'info',
									//   message: `action: ${ action }`
									// });
								}
							});
							// alert("登录成功");
						} else {
							// alert("用户名与密码不匹配");
							that.$alert('用户名与密码不匹配', '警告', {
								confirmButtonText: '确定',
								// callback: action => {
									// window.location.href = "index.html";
									// that.$message({
									//   type: 'info',
									//   message: `action: ${ action }`
									// });
								// }
							});
						}
					})
				} else {
					console.log('error submit!!');
					that.$alert('用户名和密码不能为空', '警告', {
						confirmButtonText: '确定',
						callback: action => {
							// this.$message({
							//   type: 'info',
							//   message: `action: ${ action }`
							// });
						}
					});
					return false;
				}
			});
		},
		resetForm(formName) {
			this.$refs[formName].resetFields();
		}
	}
})

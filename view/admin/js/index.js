new Vue({
	el: '#app',
	data() {
		return {
			// activeName: 'second',
			activeName: 'first',
			ticketList: [],
			myOrderList: [],
			gridData: [],
			index: 1,
			size: 10,
			total: 0,
			dialogTableVisible: false,
			dialogFormVisible: false,
			form: {
				title: '',
				details : '',
				price : '',
				name: '',
				region: '',
				date1: '',
				date2: '',
				delivery: false,
				type: [],
				resource: '',
				desc: ''
			},
			formLabelWidth: '120px'
		};
	},
	methods: {
		handleEdit(index, row) {
			// alert(index, row);
		},
		handleClick(tab, event) {
			if(tab.index == 0){
				this.handleGetTicketList();
			} else if(tab.index == 1){
				this.handleGetOrderList()
			}
		},
		handleExit() {
			location.href = "login.html"
		},
		handleDeleteOrder(index, row) {
			var that = this;
			that.$confirm('此操作将永久删除该门票类型, 是否继续?', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				console.log(row)
				// console.log(baseUrl+"/v1/order/deleteOrderById?u_id="+id+"&id="+row.row.id)
				axios.delete(baseUrl+"/v1/ticket/delTicketList",{
					params: {
						id: row.id,
					}
				}).then(function(res){
					if(res.status==200 && res.data.code == 200){
						that.$message({
							type: 'success',
							message: '操作成功!'
						});
						that.handleGetTicketList();
					} else {
						that.$message({
							type: 'success',
							message: '操作失败!'
						});
					}
				
				})
			}).catch((err) => {
				that.$message({
					type: 'info',
					message: '已取消删除'
				});
				console.log(err)
			});
		},
		handleAddOrder(){
			var that = this;
			that.$confirm('请确认是否提交?', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				axios.post(baseUrl+"/v1/ticket/addTicket", Qs.stringify(
					{
						title: that.form.title,
						details: that.form.details,
						price: that.form.price,
					}),{
					headers: {
						"Content-Type": "application/x-www-form-urlencoded",
					}
				}).then(function(res){
					if(res.status==200 && res.data.code == 200){
						that.$message({
							type: 'success',
							message: '操作成功!'
						});
						that.handleGetTicketList()
						that.handleOpenAddTicketFrm(false)
					} else {
						that.$message({
							type: 'success',
							message: '操作失败!'
						});
					}
				
				})
			}).catch((err) => {
				that.$message({
					type: 'info',
					message: '已取消购买'
				});
				console.log(err)
			});
		},
		handleGetOrderList(){
			this.getOrderList(1)
		},
		handleGetTicketList(){
			var that = this;
			axios.get(baseUrl+"/v1/ticket/ticketList").then(function(res) {
				that.ticketList = res.data.data;
			})
		},
		getOrderList(index){
			var that = this;
			that.index = index;
			axios.get(baseUrl+"/v1/order/orderList?index="+that.index+"&size="+that.size).then(function(res) {
				// var data = res.data;
				var ctx = res.data;
				that.index = ctx.data.index;
				that.size = ctx.data.size
				that.total = ctx.data.total;
				that.myOrderList = ctx.data.list;
				console.log(ctx)
				// console.log(res.data)
			})
		},
		handleOpenAddTicketFrm(visible){
			this.dialogFormVisible = visible;
			if(visible){
				this.form.title = '';
				this.form.details = '';
				this.form.price = '';
			}
		}
	},
	mounted: function() {
		this.handleGetOrderList()
		this.handleGetTicketList()
	}
})

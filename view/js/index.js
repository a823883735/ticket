new Vue({
	el: '#app',
	data() {
		return {
			activeName: 'second',
			tableData: [{
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"

			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}, {
				title: "海豚表演",
				time: "2021/1/2-2021/1-5",
				price: "999"
			}],
			gridData: [{
				date: '2016-05-02',
				name: '王小虎',
				address: '上海市普陀区金沙江路 1518 弄'
			}, {
				date: '2016-05-04',
				name: '王小虎',
				address: '上海市普陀区金沙江路 1518 弄'
			}, {
				date: '2016-05-01',
				name: '王小虎',
				address: '上海市普陀区金沙江路 1518 弄'
			}, {
				date: '2016-05-03',
				name: '王小虎',
				address: '上海市普陀区金沙江路 1518 弄'
			}],
			dialogTableVisible: false,
			dialogFormVisible: false,
			form: {
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
		handleClick() {

		}
	},
	mounted: function() {
		axios.get("http://106.13.237.16:10000/v1/order/myOrderList?uid=" + localStorage.getItem("id") +
			"&index=1&size=10").then(function(res) {
			// var data = res.data;
			// console.log(res.data)
		})
		axios.get("http://localhost:10000/v1/ticket/ticketList").then(function(res) {
			// var data = res.data;
			// console.log(res.data)
		})
	}
})

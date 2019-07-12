$(function(){
	$("#add_list").click(function(){
		$("#add_interface").toggle(300)
	})
	$("#add_card").bind("input propertychange",function sft() {
        var reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/; 
		if(reg.test($("#add_card").val()) === false)  {
		    $("#yz").css("visibility","visible")
		    return false
		}else{
			$("#yz").css("visibility","hidden")
			return true
		}
    });
	$("#add_sub").bind("click",function(){
		
		if($("#yz").css("visibility")==="visible"){
			alert("您提交的身份证信息不合法")
			return false
		}else{

			var name=$("#add_name").val()
			var sfz =$("#add_card").val()
			console.log($("#add_card").val())
			$.ajax({
				type:"Get",
				url:"http://localhost:19999/addblacklist",
				datatype:'json',
				async:false,											//添加黑名单
				data:{
					'name':name,
					'sfz':sfz
				},
				success:function(data){
					event.preventDefault()
					var hell = $("<div></div>")
					var nrong = $("<div></div>")
					var anni =$("<div></div>")
					$('#cont').append(hell)
					hell.addClass("panel panel-default col-lg-3")
					hell.append(nrong)
					nrong.addClass("panel-body wd")
					nrong.append($("#add_name").val()+"<a>&nbsp&nbsp</a>:<a>&nbsp&nbsp</a>"+$("#add_card").val())
					var dele=$("<input type='button' value='删除'>")
					dele.addClass("navbar-right btn btn-success btn-sm delete")
					nrong.append(dele)
					
					var his = $("<li></li>")
					his.addClass("list-group-item list-group-item-success")
					his.append($("#add_name").val()+"<a>&nbsp&nbsp</a>:<a>&nbsp&nbsp</a>"+$("#add_card").val())
					$(".his_ul").append(his)
				},
				error:function(){
					alert('9999')
				}
				
			});
		}
	})
//	$("#lookup").bind("click",function(){
//
//			console.log($("#look").val())
//			var look=$("#look").val()
//			$.ajax({
//				type:"Get",
//				url:"http://localhost:19999/addblacklist",
//				datatype:'json',
//				async:false,
//				data:{
//					'look':look													//查找
//				},
//				success:function(data){
//					event.preventDefault()
//				},
//				error:function(){
//					alert('9999')
//				}
//				
//			});
//		
//	})
	$("div").on("click",".delete",function(){
		console.log(111)
		$(this).parent().parent().hide(200)
	})
	$("#history").bind("click",function(){
		$("#history_infor").toggle(200)
	})
})

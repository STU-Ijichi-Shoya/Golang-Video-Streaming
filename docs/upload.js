{
    // let droparea=document.getElementById("droparea")
    let selectbtn=document.getElementById("videoFile")
    let title=document.getElementById("title")
    let userName=document.getElementById("user-name")

    let videoPreviewContainer=document.getElementById('video-preview-container')
    let videoPreviewElement=document.getElementById('video-preview-element');

    let fileState=false;
    selectbtn.addEventListener("change", ()=>{
        const file=selectbtn.files[0];
        if(file.size>1024*1024*200){
            alert('動画は200MB以下にしてください')
            file.value='';
            return 
        }
        fileState=true;
        
        videoPreviewContainer.hidden=false;
        let blobUrl=window.URL.createObjectURL(file)
        videoPreviewElement.src=blobUrl;
        
    })
    function remove(s){
        return s.split(" ").join("").split("\t").join("");
    }
    function resetBtn(){
        if(fileState){
            fileState=false
            selectbtn.files[0].value=''
        }
        title.value=''
        userName.value=''
        videoPreviewElement.src=''
        videoPreviewContainer.hidden=true;
    }
    function uploadBtn(){
        if(!fileState){
            alert("動画を選択してください")
            return
        }
        const titlestr=title.value;
        let r=remove(titlestr)
        if(r.length<=0){
            alert("タイトルを入力してください")
            return
        }
        let name=userName.value;

        if(remove(name).length==0){
            name="名無しさん"
        }


        const formData=new FormData();

        formData.append("inputText",r);
        formData.append("videoFile",selectbtn.files[0]);
        const param = {
            method: "POST",
            body: formData
        }
        fetch("/upload",param).then((res)=>{
            const total = res.headers.get('content-length');
            document.getElementById("uploadProgress").style.display="block"
            let reader = res.body.getReader();
            let chunk = 0;
            const progressBar=document.getElementById("uploadProgress")
            reader.read().then(function processResult(result){
                    if(result.done){
                        return console.log('UPLOAD OK')
                    }
                    chunk += result.value.length;
                    // log(`received: ${chunk}(${Math.round(chunk/total * 100)} %)`);
                    let progress = Math.round(chunk / total * 100)
                    progressBar.setAttribute('value',progress)
                    // 再帰する
                    return reader.read().then(processResult);
                }
            )

        }).then((json)=>{
            alert("アップロードに成功しました")
            alert("続けてアップロードする場合は選択してください")
            resetBtn()
        }).catch((error)=>{
            alert("エラーが発生しました，時間を置くか，違う動画をアップロードしてください")

        })
        console.log(name,titlestr)

    }
}
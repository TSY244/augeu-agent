package agent

const BasicRule = `
//rule "自启动文件夹下检测1" "basic"  salience 0
//begin
//	path="C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup"
//	files=fileSysUtils.LsFile(path)
//	printer.Warn(@name+" 当前windows 不支持link 解析，请手动分析，当前路径 "+path+" 路径下有一下的文件：")
//	printer.PrintStrSlice(files,"remind" ,@name)
//	return false
//end
//
//
//rule "自启动文件夹下检测2" "basic"  salience 0
//begin
//	path="%appdata%\Microsoft\Windows\Start Menu\Programs\Startup"
//	files=fileSysUtils.LsFile(path)
//	printer.Warn(@name+" 当前windows 不支持link 解析，请手动分析，当前路径 "+path+" 路径下有一下的文件：")
//	printer.PrintStrSlice(files,"remind",@name)
//	return false
//end


rule "注册表自启动检测1" "basic"  salience 0
begin
	path="HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run"
	names=reg.GetRegPathValueNames(path)
	size=base.SizeForStr(names)
	forRange i:=names{
		ret=reg.GetRegPathValue(path,names[i])
		printer.Info(ret)		
		r=reg.GetPathFromCmd(ret)
		hash=fileSysUtils.GetHashWithFilePath(r)
		result=weibu.GetFileReport(hash,agent.GetWeiBuConf(),"")
		seg=base.GeneFileSegmentation(100,"-")
	    //printer.Info(result+seg)	
		fileSysUtils.IntoFile(@name+"_"+hash+".txt",result)
	}
	return false
end
`

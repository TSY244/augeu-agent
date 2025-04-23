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
//
//
//rule "注册表自启动检测1" "basic"  salience 0
//begin
//	path="HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run"
//	names=reg.GetRegPathValueNames(path)
//	size=base.SizeForStr(names)
//	flag=true
//	forRange i:=names{
//		ret=reg.GetRegPathValue(path,names[i])
//		printer.Info(ret,@name)		
//		r=reg.GetPathFromCmd(ret)
//		hash=fileSysUtils.GetHashWithFilePath(r)
//		result=weibu.GetFileReport(hash,agent.GetWeiBuConf(),"")
//		seg=base.GeneFileSegmentation(100,"-")
//	    //printer.Info(result+seg)	
//		flag=flag&&check.CheckHash(hash,agent.GetWeiBuConf(),"",0.5)
//		fileSysUtils.IntoFile(@name+"_"+hash+".txt",result)
//	}
//	return flag
//end
//
//rule "注册表自启动检测2" "basic"  salience 0
//begin
//	path="HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\RunOnce"
//	names=reg.GetRegPathValueNames(path)
//	size=base.SizeForStr(names)
//	flag=true
//	forRange i:=names{
//		ret=reg.GetRegPathValue(path,names[i])
//		printer.Info(ret,@name)		
//		r=reg.GetPathFromCmd(ret)
//		hash=fileSysUtils.GetHashWithFilePath(r)
//		result=weibu.GetFileReport(hash,agent.GetWeiBuConf(),"")
//		seg=base.GeneFileSegmentation(100,"-")
//	    //printer.Info(result+seg)	
//		flag=flag&&check.CheckHash(hash,agent.GetWeiBuConf(),"",0.5)
//		fileSysUtils.IntoFile(@name+"_"+hash+".txt",result)
//	}
//	return flag
//end
//
//rule "LogonScript（优先于AV）" "basic"  salience 0
//begin
//	path="HKEY_CURRENT_USER\Environment"
//	names=reg.GetRegPathValueNames(path)
//	flag=true
//	flag2=strUtils.StrSliceContains(names,"UserInitMprLogonScript")
//	if flag2{
//		printer.Warn(@name+" 中存在UserInitMprLogonScript 键在HKEY_CURRENT_USER\Environment 下",@name)
//	}else{
//		printer.Info(@name+" 中不存在UserInitMprLogonScript 键在HKEY_CURRENT_USER\Environment 下，所以安全",@name)
//		return flag
//	}
//	value:=reg.GetRegPathValue(path,"UserInitMprLogonScript")
//	hashRet=fileSysUtils.GetHashWithFilePath(value)
//	printer.Info(hashRet,@name)
//	flag=check.CheckHash(hashRet,agent.GetWeiBuConf(),"",0.5)
//	return flag
//end
//
//rule "屏幕保护后门检测" "basic"  salience 0
//begin
//	path="HKEY_CURRENT_USER\Control Panel\Desktop"
//	names=reg.GetRegPathValueNames(path)
//	flag=true
//	flag2=strUtils.StrSliceContains(names,"SCRNSAVE.EXE")
//	if flag2{
//		printer.Warn(@name+" 中存在SCRNSAVE.EXE 键在"+path+" 下",@name)
//	}else{
//		printer.Info(@name+" 中不存在SCRNSAVE.EXE 键在"+path+" 下，所以安全",@name)
//		return flag
//	}
//	value:=reg.GetRegPathValue(path,"SCRNSAVE.EXE")
//	hashRet=fileSysUtils.GetHashWithFilePath(value)
//	printer.Info(hashRet,@name)
//	flag=check.CheckHash(hashRet,agent.GetWeiBuConf(),"",0.5)
//	return flag
//end
//
//
//rule "inf文件后门（先于大部分软件）" "basic"  salience 0
//begin
//	path="HKEY_CURRENT_USER\Software\Microsoft\IEAK\GroupPolicy\PendingGPOs"
//	flag=reg.IsHavePath(path)
//	if flag{
//		printer.Warn(@name+" 中存在HKEY_CURRENT_USER\Software\Microsoft\IEAK\GroupPolicy\PendingGPOs 路径",@name)
//	}else{
//		printer.Info(@name+" 中不存在HKEY_CURRENT_USER\Software\Microsoft\IEAK\GroupPolicy\PendingGPOs 路径，所以安全",@name)
//		return true
//	}
//	exePath=reg.GetRegPathValue(path,"Path1")
//	printer.Info(exePath,@name)
//	hashRet=fileSysUtils.GetHashWithFilePath(exePath)
//	printer.Info(hashRet,@name)
//	flag=check.CheckHash(hashRet,agent.GetWeiBuConf(),"",0.5)
//	return flag
//end
//
//
//rule "系统级自启动-注册表统一检测任务" "basic"  salience 0
//begin
//	pathes=strUtils.CraterStrSlice("HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Run",
// "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnce",
// "HKEY_LOCAL_MACHINE\SOFTWAREWOW6432Node\Microsoft Windows\CurrentVersion\Run",
// "HKEY_LOCAL_MACHINE\SOFTWAREWOW6432Node\Microsoft\Windows\CurrentVersion\RunOnce",
// "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnceEx\0001",
// "HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\RunOnceEx\0001\Depend")
//
//	forRange i:=pathes{
//		path=pathes[i]
//		isHavePath=reg.IsHavePath(path)
//		if !isHavePath{
//			printer.Info(@name+" 中不存在"+path+" 路径，所以安全",@name)
//			continue
//		}
//		names=reg.GetRegPathValueNames(path)
//		size=base.SizeForStr(names)
//		flag=true
//		forRange i:=names{
//			ret=reg.GetRegPathValue(path,names[i])
//			printer.Info(ret,@name)
//			r=reg.GetPathFromCmd(ret)
//			hash=fileSysUtils.GetHashWithFilePath(r)
//			result=weibu.GetFileReport(hash,agent.GetWeiBuConf(),"")
//			seg=base.GeneFileSegmentation(100,"-")
//		    //printer.Info(result+seg)
//		}
//	}
//	return false
//end
//
//
//rule "镜像劫持" "basic"  salience 0
//begin
//	path="HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Image File Execution Options"
//	isHavePath=reg.IsHavePath(path)
//	if !isHavePath{
//		printer.Info(@name+" 中不存在"+path+" 路径，所以安全",@name)
//		return true
//	}
//	subKeys=reg.GetPathSubKeys(path)
//	pathWithSubKey=strUtils.AddPrefixs(subKeys,path+strUtils.GeneABackslash())
//	flag=false
//	forRange i:=pathWithSubKey{
//		newPath=pathWithSubKey[i]
//		printer.Info(newPath)
//		names=reg.GetRegPathValueNames(newPath)
//		flag2=strUtils.StrSliceContainsIgnoreCase(names,"debugger")
//		if flag2{
//			printer.Warn(@name+" 中存在debugger 键在"+path+" 下",@name)
//		}else{
//			//printer.Info(@name+" 中不存在debugger 键在"+path+" 下，所以安全",@name)
//			continue
//		}
//		ret=reg.GetRegPathValue(newPath,"debugger")
//		printer.Info(ret,@name)
//		r=reg.GetPathFromCmd(ret)
//		printer.Warn("存在恶意键debugger: "+r,@name)
//		hash=fileSysUtils.GetHashWithFilePath(r)
//		result=check.CheckHash(hash,agent.GetWeiBuConf(),"",0.5)
//		flag=flag && !result
//	}
//	return flag
//end
	
//
//rule "隐藏用户" "basic"  salience 0
//begin
//	path="HKEY_LOCAL_MACHINE\SAM\SAM\Domains\Account\Users\Names"
//	subKeys=reg.GetPathSubKeys(path)
//	target=strUtils.GetDollarSign()
//	flag2=strUtils.StrSliceContainsIgnoreCase(subKeys,target)
//	if flag2{
//		printer.Warn(@name+" 中存在$ 键在"+path+" 下，可能存在隐藏用户",@name)
//		return false
//	}
//	pathWithSubKey=strUtils.AddPrefixs(subKeys,path+strUtils.GeneABackslash())
//	flag=0
//	forRange i:=pathWithSubKey{
//		newPath=pathWithSubKey[i]
//		printer.Info(newPath)
//		names=reg.GetRegPathValueNames(newPath)
//		ret=reg.GetDefaultRegPathValue(newPath)
//		if ret=="500"{
//			flag=flag+1
//		}
//	}
//	return flag<2
//end

//
//rule "AppCertDlls" "basic" salience 0
//begin
//	path="HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\SessionManager\AppCertDlls"
//	isHavePath=reg.IsHavePath(path)
//	if !isHavePath{
//		printer.Info(@name+" 中不存在"+path+" 路径，所以安全",@name)
//		return true
//	}
//	printer.Warn(@name+" 存在"+path+" 路径",@name)
//	ret=reg.GetDefaultRegPathValue(path)
//	if ret ==""{
//	printer.Info(@name+"没有键值",@name)
//		return false
//	}
//end


rule "AppInit_DLLs " "basic" salience 0
begin
	path="HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Windows\AppInit_DLLs"
	isHavePath=reg.IsHavePath(path)
	if !isHavePath{
		printer.Info(@name+" 中不存在"+path+" 路径，所以安全",@name)
		return true
	}
	printer.Warn(@name+" 存在"+path+" 路径",@name)
	ret=reg.GetDefaultRegPathValue(path)
	if ret ==""{
	printer.Info(@name+"没有键值",@name)
		return false
	}
end
`

require 'dicom'
require 'faraday'
include DICOM

def put (url,path,dcm) # 发送数据
    conn = Faraday.new(:url => url) do |faraday|
        faraday.request  :url_encoded             # form-encode POST params
        faraday.response :logger                  # log requests to STDOUT
        faraday.adapter  Faraday.default_adapter  # make requests with Net::HTTP
    end
    conn.put do |req|
        req.url path
        req.headers['Content-Type'] = 'application/json'
        req.body = dcm.to_json
    end
end

def do_it file_path # 处理
    # Read file:
    dcm = DObject.read(file_path)
    begin
        fd_obj = dcm.send("sop_instance_uid")
        puts "*****     sopID为：" << fd_obj.value() << "   *****"
        str = put('http://localhost:9200',"/dicom/external/#{fd_obj.value()}",dcm)
        puts ">>>>>>  结果 s  <<<<<<"
        p str.body.to_s
        puts ">>>>>>  结果 e  <<<<<<"
        puts "****  处理完成  ****"
    rescue => exception
        puts "****  错误  ****"
        puts exception
    end
end

def traverse_dir(file_path)# 遍历文件
    if File.directory? file_path
        Dir.foreach(file_path) do |file|
            if file !="." and file !=".."
                traverse_dir(file_path+"/"+file)
            end
        end
    else
        puts "--------------   正在处理文件: #{File.basename(file_path)}, Size:#{File.size(file_path)}   -----------------"
        # if File.basename(file_path).split('.').last.downcase == "dcm"
            do_it(file_path)
        # end
    end
end

# 走你
traverse_dir('/Users/Orange/Desktop/ruby_dicom_test/dicoms')


import os

from GoWriter import GoWriter
from htmlParser import ApiCollector

needToBuild = []
datafiles = []
for root, dirs, files in os.walk("./Data", topdown=True):
    # create the directory
    for dir in dirs:
        if not os.path.exists("./structures/" + dir):
            os.mkdir("./structures/" + dir)

    for file in files:
        needToBuild.append(os.path.join(root, file).replace("Data", "structures"))
        datafiles.append(os.path.join(root, file))

i = 0
while i < len(needToBuild):
    print("Parsing response : " + needToBuild[i])
    parser = ApiCollector(datafiles[i])
    # get the file name without extension
    parser.ParseResponse(os.path.splitext(os.path.basename(datafiles[i]))[0])
    GoWriter(needToBuild[i]).write(parser.post_process())
    i += 1

# //*[@id="parameters_block"]/div/table/tbody/tr[2]/td[1]/div/label/span[1]
# //*[@id="parameters_block"]/div/table/tbody/tr[3]/td[1]/div/label/span[1]
# //*[@id="parameters_block"]/div/table/tbody/tr[4]/td[1]/div/label/span[1]
# //*[@id="parameters_block"]/div/table/tbody/tr[5]/td[1]/div/label/span[1]


# //*[@id="Parameters-field-account_id"]
# //*[@id="Parameters-field-fields"]
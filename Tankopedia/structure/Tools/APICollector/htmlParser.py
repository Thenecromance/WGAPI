from lxml import etree

from Parameter import Parameter, ParameterFile


class ApiCollector:
    html = None
    all_args = []
    pFile = []

    def __init__(self, file_path):
        self.html = etree.parse(file_path, etree.HTMLParser())
        self.all_args = []
        self.pFile = []
        pass

    def get_args_count(self):
        return len(self.html.xpath('//*[@id="response_block"]/div/table/tbody')[0])

    def get_group(self, idx):
        return self.html.xpath('//*[@id="response_block"]/div/table/tbody/tr[' + str(idx) + ']')[0]

    def parse_Field(self, group) -> str:
        level_count = len(group.xpath('./td[1]')[0])
        Field = ""
        if level_count == 1:
            Field = \
                group.xpath('./td[1]/span/span')[
                    0].text
        else:
            i = 1
            while i <= level_count:
                Field += group.xpath('./td[1]/span[' + str(i) + ']/span[1]')[0].text + "."
                i += 1
            Field = Field[:-1]
        return Field

    def parse_Type(self, group) -> str:
        # //*[@id="response_block"]/div/table/tbody/tr[828]/td[2]
        Field = ""
        try:
            Field = group.xpath('./td[2]')[0].text
        except:
            return ""
        return Field

    def parse_Description(self, group) -> str:
        # //*[@id="response_block"]/div/table/tbody/tr[828]/td[3]/div/p
        try:
            Field = group.xpath('./td[3]/div/p')[0].text
        except:
            return ""
        return Field

    def ParseResponse(self, default_name: str):
        idx = 2
        while idx <= self.get_args_count():
            group = self.get_group(idx)
            param = None
            # print(len(group.xpath('./td[1]/span')))
            if len(group.xpath('./td[1]/span')) >= 1:
                param = Parameter(
                    self.parse_Field(group),
                    self.parse_Type(group),
                    self.parse_Description(group))
                self.all_args.append(param)
            idx += 1

        idx = 0
        # self.pFile
        file = ParameterFile(default_name, "")
        while idx < len(self.all_args):
            if self.all_args[idx].Type is None:
                # print("New " + self.all_args[idx].Field)
                self.pFile.append(file)
                file = ParameterFile(self.all_args[idx].Field, self.all_args[idx].Description)
                idx += 1
            else:
                file.append(self.all_args[idx])
                idx += 1
        self.pFile.append(file)

        return self.pFile

    def post_process(self):
        root: ParameterFile = self.pFile[0]
        relations = {root.file_name: root}
        # generate each files relationship
        for file in self.pFile[1:]:
            relations[file.file_name] = file
            if file.file_name.find(".") == -1:
                root.children[file.file_name] = file
            else:
                # print(file.file_name[:file.file_name.rfind(".")])
                parentName = file.file_name[:file.file_name.rfind(".")]
                relations[parentName].children[file.file_name] = file
                # root.children[file]

        # rename each node's name
        for file in self.pFile[1:]:
            if file.file_name.find(".") != -1:
                file.file_name = file.file_name[file.file_name.rfind(".") + 1:]
        # self.print_child(root)

        return root

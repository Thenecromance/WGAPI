class Parameter:
    """Parameter class"""

    # Field: str = ""
    # Type: str = ""
    # Description: str = ""

    def __init__(self, field, type, desc):
        """Parameter constructor"""
        self.Field = field
        self.Type = type
        self.Description = desc

    def __str__(self):
        """Parameter string representation"""
        return self.Field + ": " + str(self.Type) + ":" + self.Description

    # from snake case to camel case
    def camel_case(self) -> str:
        return "".join(x.capitalize() or "_" for x in self.Field.split(".")[-1].split("_"))

    def type_to_go(self):
        if self.Type == "numeric":
            return "int32" + "/*" + self.Type + "*/"
        elif self.Type == "float":
            return "float64" + "/*" + self.Type + "*/"
        elif self.Type == "timestamp":
            return "int64" + "/*" + self.Type + "*/"
        elif self.Type == "list of integers":
            return "[]int32" + "/*" + self.Type + "*/"
        elif self.Type == "list of strings":
            return "[]string" + "/*" + self.Type + "*/"
        elif self.Type == "associative array":
            return "map[string]string" + "/*" + self.Type + "*/"
        elif self.Type == "boolean":
            return "bool" + "/*" + self.Type + "*/"
        else:
            return self.Type + "/*" + self.Type + "*/"

    def add_tag(self):
        result = "`"
        result += "json:\"" + self.Field.split(".")[-1] + "\" "
        result += "xml:\"" + self.Field.split(".")[-1] + "\" "
        result += "`"

        return result
        # return "`json:\"" + self.Field + "\"`"

    def clean_description(self):
        return self.Description.replace("\n", "").replace("\r", " ").replace("\t", " ") + "\n"

    def write_to_file(self, f):
        f.write(self.camel_case())
        f.write(" " + self.type_to_go())
        f.write(" " + self.add_tag())
        f.write(" // " + self.clean_description())
        f.write("\n")


class ParameterFile:
    """ParameterFile class"""

    def __init__(self, file_name, description):
        """ParameterFile constructor"""
        self.children = {}
        self.Parameters = []
        self.file_name = file_name
        self.Description = description

    def append(self, parameter):
        """Append Parameter to ParameterFile"""
        self.Parameters.append(parameter)

    # from snake case to camel case
    def camel_case(self) -> str:
        return "".join(x.capitalize() or "_" for x in self.file_name.split("_"))

    def add_tag(self):
        result = "`"
        result += "json:\"" + self.file_name + "\" "
        result += "xml:\"" + self.file_name + "\" "
        result += "`"
        return result

    def clean_description(self):
        return self.Description.replace("\n", "\n//").replace("\r", " ").replace("\t", " ")

    def write_as_param(self, f):
        f.write(self.camel_case())
        f.write(" *" + self.camel_case() + self.add_tag())
        f.write(" // " + self.clean_description())
        f.write("\n")

    def __str__(self):
        """ParameterFile string representation"""
        return str(self.Parameters)

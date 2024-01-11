import os

from Parameter import ParameterFile, Parameter
import os


class GoWriter:

    def __init__(self, dir):

        if dir.rfind(".") != -1:
            self.sub_dir = dir[:dir.rfind(".")]
        else:
            self.sub_dir = dir
        if not os.path.exists(self.sub_dir):
            os.mkdir(self.sub_dir)

        pass

    def _write_param(self, f, param: list[Parameter]):
        for p in param:
            p.write_to_file(f)
            pass

    def _write_children(self, f, children: dict[str, ParameterFile]):
        for child in children.values():
            child.write_as_param(f)
            pass

    def write(self, file: ParameterFile):
        with open(self.sub_dir + "/" + file.camel_case() + ".go", 'w', encoding="utf-8") as f:
            f.write("package structure\n")
            f.write("type " + file.camel_case() + " struct {\n")
            self._write_param(f, file.Parameters)
            self._write_children(f, file.children)
            f.write("}")

            for child in file.children.values():
                self.write(child)
                pass
